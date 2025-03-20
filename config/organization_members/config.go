package organizationmembers

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_organization_members resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_members", func(r *config.Resource) {
		r.ShortGroup = "organizationmembers"
	})
}
