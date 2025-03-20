package pages

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_pages resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_pages", func(r *config.Resource) {
		r.ShortGroup = "pages"
	})
}
