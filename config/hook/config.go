package hook

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_hook", func(r *config.Resource) {
		r.ShortGroup = "hook"
	})
}
