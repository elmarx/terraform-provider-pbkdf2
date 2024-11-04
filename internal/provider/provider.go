// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure Pbkdf2Provider satisfies various provider interfaces.
var _ provider.Provider = &Pbkdf2Provider{}
var _ provider.ProviderWithFunctions = &Pbkdf2Provider{}

// Pbkdf2Provider defines the provider implementation.
type Pbkdf2Provider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Pbkdf2ProviderModel describes the provider data model.
type Pbkdf2ProviderModel struct {
}

func (p *Pbkdf2Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pbkdf2"
	resp.Version = p.version
}

func (p *Pbkdf2Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *Pbkdf2Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data Pbkdf2ProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *Pbkdf2Provider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *Pbkdf2Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *Pbkdf2Provider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewPbkdf2Sha512Function,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Pbkdf2Provider{
			version: version,
		}
	}
}
