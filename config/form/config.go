package form

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_form", func(r *config.Resource) {
		r.ShortGroup = "form"
	})
}
