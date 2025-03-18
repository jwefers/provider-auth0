package organizationmemberrole

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_member_role", func(r *config.Resource) {
		r.ShortGroup = "organizationmemberrole"
	})
}
