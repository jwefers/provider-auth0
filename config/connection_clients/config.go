package connectionclients

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_clients", func(r *config.Resource) {
		r.ShortGroup = "connectionclients"
	})
}
