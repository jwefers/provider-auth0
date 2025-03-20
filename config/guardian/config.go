package guardian

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_guardian resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_guardian", func(r *config.Resource) {
		r.ShortGroup = "guardian"
	})
}
