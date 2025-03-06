package tenant

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_tenant", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "tenant"
	})
}
