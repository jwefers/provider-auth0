package userroles

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_roles", func(r *config.Resource) {
		r.ShortGroup = "userroles"
	})
}
