package branding

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_branding", func(r *config.Resource) {
		r.ShortGroup = "branding"
	})
}
