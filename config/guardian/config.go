package guardian

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_guardian in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_guardian", func(r *config.Resource) {
		r.ShortGroup = "guardian"
	})
}
