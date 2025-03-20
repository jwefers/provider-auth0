package resourceserverscope

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_resource_server_scope in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_resource_server_scope", func(r *config.Resource) {
		r.ShortGroup = "resourceserverscope"
	})
}
