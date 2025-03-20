package customdomain

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_custom_domain resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_custom_domain", func(r *config.Resource) {
		r.ShortGroup = "customdomain"
	})
}
