package clientcredentials

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_client_credentials resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client_credentials", func(r *config.Resource) {
		r.ShortGroup = "clientcredentials"
	})
}
