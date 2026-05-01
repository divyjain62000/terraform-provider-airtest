package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// ✅ Correct type
var testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"demo": providerserver.NewProtocol6WithError(New()),
}
