package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	customdatasource "terraform-provider-customazure/datasource"
	"terraform-provider-customazure/client"
)


type CustomAzureProvider struct{}

func New() provider.Provider {
	return &CustomAzureProvider{}
}

func (p *CustomAzureProvider) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "customazure"
}

func (p *CustomAzureProvider) Schema(
	_ context.Context,
	_ provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"tenant_id": schema.StringAttribute{
				Required: true,
			},
			"client_id": schema.StringAttribute{
				Required: true,
			},
			"***REMOVED***": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *CustomAzureProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var config struct {
		TenantID     string `tfsdk:"tenant_id"`
		ClientID     string `tfsdk:"client_id"`
		ClientSecret string `tfsdk:"***REMOVED***"`
	}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

graphClient, err := client.NewGraphClient(
	config.TenantID,
	config.ClientID,
	config.ClientSecret,
)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to authenticate to Microsoft Graph",
			err.Error(),
		)
		return
	}

resp.DataSourceData = graphClient

}

func (p *CustomAzureProvider) DataSources(
	_ context.Context,
) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		customdatasource.NewUsersDataSource,
	}
}

func (p *CustomAzureProvider) Resources(
	_ context.Context,
) []func() resource.Resource {
	return []func() resource.Resource{}
}
