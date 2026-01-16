package datasource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-customazure/client"
)

type UsersDataSource struct {
	client *client.GraphClient
}

func NewUsersDataSource() datasource.DataSource {
	return &UsersDataSource{}
}

func (d *UsersDataSource) Metadata(
	_ context.Context,
	_ datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = "customazure_users"
}

func (d *UsersDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"users": schema.ListAttribute{
				Computed: true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *UsersDataSource) Configure(
	_ context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	if req.ProviderData != nil {
		d.client = req.ProviderData.(*client.GraphClient)
	}
}

func (d *UsersDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	users, err := d.client.Client.Users().Get(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read users", err.Error())
		return
	}

	var userNames []string
	for _, user := range users.GetValue() {
		userNames = append(userNames, *user.GetDisplayName())
	}

	resp.State.Set(ctx, map[string]interface{}{
		"id":    "users",
		"users": userNames,
	})
}
