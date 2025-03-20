package connectionclient

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_connection_client in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_client", func(r *config.Resource) {
		r.ShortGroup = "connectionclient"
	})
}
