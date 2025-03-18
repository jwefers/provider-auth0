package role

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_role", func(r *config.Resource) {
		r.ShortGroup = "role"
	})
}
