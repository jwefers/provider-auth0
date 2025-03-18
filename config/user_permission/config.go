package userpermission

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_permission", func(r *config.Resource) {
		r.ShortGroup = "userpermission"
	})
}
