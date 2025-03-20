package clientcredentials

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_client_credentials in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_client_credentials", func(r *config.Resource) {
		r.ShortGroup = "clientcredentials"
	})
}
