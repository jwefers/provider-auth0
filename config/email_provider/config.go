package emailprovider

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_email_provider in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_email_provider", func(r *config.Resource) {
		r.ShortGroup = "emailprovider"
	})
}
