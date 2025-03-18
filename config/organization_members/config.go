package organizationmembers

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_members", func(r *config.Resource) {
		r.ShortGroup = "organizationmembers"
	})
}
