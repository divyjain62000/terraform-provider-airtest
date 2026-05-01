package resources

import (
	"testing"

	"github.com/airtest/terraform-provider-airtest/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// ✅ ADD THIS
var testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"airtest": providerserver.NewProtocol6WithError(provider.New()),
}

func TestAccUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
provider "airtest" {
  api_token = "test123"
  api_url   = "https://api.example.com"
}

resource "airtest_user" "test" {
  name  = "John Doe"
  email = "john@example.com"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("airtest_user.test", "name", "John Doe"),
					resource.TestCheckResourceAttr("airtest_user.test", "email", "john@example.com"),
				),
			},
		},
	})
}
