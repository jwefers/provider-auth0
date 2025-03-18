package organization

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization", func(r *config.Resource) {
		r.ShortGroup = "organization"
	})
}
