package rolepermission

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_role_permission resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_role_permission", func(r *config.Resource) {
		r.ShortGroup = "rolepermission"
	})
}
