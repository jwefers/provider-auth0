package organizationconnection

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_connection", func(r *config.Resource) {
		r.ShortGroup = "organizationconnection"
	})
}
