package user

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_user in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user", func(r *config.Resource) {
		r.ShortGroup = "user"
	})
}
