package connectionclients

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_connection_clients resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_clients", func(r *config.Resource) {
		r.ShortGroup = "connectionclients"
	})
}
