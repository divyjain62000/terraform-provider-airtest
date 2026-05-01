package main

import (
	"context"

	"github.com/airtest/terraform-provider-airtest/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/airtest/airtest", //registry.terraform.io/<org>/<repo>
	})
}
