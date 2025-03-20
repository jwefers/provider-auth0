package userroles

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_user_roles resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_roles", func(r *config.Resource) {
		r.ShortGroup = "userroles"
	})
}
