package organizationconnections

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_connections", func(r *config.Resource) {
		r.ShortGroup = "organizationconnections"
	})
}
