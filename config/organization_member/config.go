package organizationmember

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_organization_member in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_member", func(r *config.Resource) {
		r.ShortGroup = "organizationmember"
	})
}
