package action

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_action", func(r *config.Resource) {
		r.ShortGroup = "action"
	})
}
