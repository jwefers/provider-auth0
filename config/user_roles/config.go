package userroles

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_user_roles in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_roles", func(r *config.Resource) {
		r.ShortGroup = "userroles"
	})
}
