package phoneprovider

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_phone_provider resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_phone_provider", func(r *config.Resource) {
		r.ShortGroup = "phoneprovider"
	})
}
