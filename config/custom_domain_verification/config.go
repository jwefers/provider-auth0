package customdomainverification

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_custom_domain_verification resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_custom_domain_verification", func(r *config.Resource) {
		r.ShortGroup = "customdomainverification"
	})
}
