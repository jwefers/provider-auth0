package customdomain

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_custom_domain in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_custom_domain", func(r *config.Resource) {
		r.ShortGroup = "customdomain"
	})
}
