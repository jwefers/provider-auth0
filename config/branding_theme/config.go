package brandingtheme

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_branding_theme", func(r *config.Resource) {
		r.ShortGroup = "brandingtheme"
	})
}
