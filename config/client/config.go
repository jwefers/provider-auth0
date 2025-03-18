package client

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client", func(r *config.Resource) {
		r.ShortGroup = "client"
	})
}
