package userpermissions

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_permissions", func(r *config.Resource) {
		r.ShortGroup = "userpermissions"
	})
}
