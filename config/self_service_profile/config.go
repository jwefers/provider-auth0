package selfserviceprofile

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_self_service_profile in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_self_service_profile", func(r *config.Resource) {
		r.ShortGroup = "selfserviceprofile"
	})
}
