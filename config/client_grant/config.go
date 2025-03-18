package clientgrant

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client_grant", func(r *config.Resource) {
		r.ShortGroup = "clientgrant"
	})
}
