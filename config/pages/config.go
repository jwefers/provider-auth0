package pages

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_pages", func(r *config.Resource) {
		r.ShortGroup = "pages"
	})
}
