package resourceserverscopes

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_resource_server_scopes in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_resource_server_scopes", func(r *config.Resource) {
		r.ShortGroup = "resourceserverscopes"
	})
}
