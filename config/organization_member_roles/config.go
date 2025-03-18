package organizationmemberroles

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_member_roles", func(r *config.Resource) {
		r.ShortGroup = "organizationmemberroles"
	})
}
