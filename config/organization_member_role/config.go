package organizationmemberrole

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_organization_member_role in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_member_role", func(r *config.Resource) {
		r.ShortGroup = "organizationmemberrole"
	})
}
