package role

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_role in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_role", func(r *config.Resource) {
		r.ShortGroup = "role"
	})
}
