package tenant

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_tenant", func(r *config.Resource) {
		r.ShortGroup = "tenant"
	})
}
