package emailprovider

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_email_provider resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_email_provider", func(r *config.Resource) {
		r.ShortGroup = "emailprovider"
	})
}
