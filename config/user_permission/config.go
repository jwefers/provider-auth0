package userpermission

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_user_permission in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_permission", func(r *config.Resource) {
		r.ShortGroup = "userpermission"
	})
}
