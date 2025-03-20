package branding

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_branding in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_branding", func(r *config.Resource) {
		r.ShortGroup = "branding"
	})
}
