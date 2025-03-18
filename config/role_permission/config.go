package rolepermission

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_role_permission", func(r *config.Resource) {
		r.ShortGroup = "rolepermission"
	})
}
