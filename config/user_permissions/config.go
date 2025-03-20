package userpermissions

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_user_permissions resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_user_permissions", func(r *config.Resource) {
		r.ShortGroup = "userpermissions"
	})
}
