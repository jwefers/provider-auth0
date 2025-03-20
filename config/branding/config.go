package branding

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_branding resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_branding", func(r *config.Resource) {
		r.ShortGroup = "branding"
	})
}
