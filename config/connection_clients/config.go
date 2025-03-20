package connectionclients

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_connection_clients in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_clients", func(r *config.Resource) {
		r.ShortGroup = "connectionclients"
	})
}
