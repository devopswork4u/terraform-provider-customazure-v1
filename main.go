package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"terraform-provider-customazure/provider"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/devopswork4u/terraform-provider-customazure-v1",
	})
}
