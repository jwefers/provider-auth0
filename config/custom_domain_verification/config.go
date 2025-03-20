package customdomainverification

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_custom_domain_verification in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_custom_domain_verification", func(r *config.Resource) {
		r.ShortGroup = "customdomainverification"
	})
}
