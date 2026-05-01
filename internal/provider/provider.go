package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type demoProvider struct{}

func New() provider.Provider {
	return &demoProvider{}
}

// Metadata
func (p *demoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "airtest"
}

// Schema
func (p *demoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				Required:    true,
				Description: "API token for authentication",
			},
			"api_url": schema.StringAttribute{
				Optional:    true,
				Description: "API endpoint for demo provider",
			},
		},
	}
}

// Configure (already correct)
func (p *demoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Later: initialize API client using api_token, api_url
}

func (p *demoProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource, // if you have one
	}
}

func (p *demoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
