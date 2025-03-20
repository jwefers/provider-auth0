package phoneprovider

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_phone_provider in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_phone_provider", func(r *config.Resource) {
		r.ShortGroup = "phoneprovider"
	})
}
