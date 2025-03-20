package selfserviceprofilecustomtext

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_self_service_profile_custom_text resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_self_service_profile_custom_text", func(r *config.Resource) {
		r.ShortGroup = "selfserviceprofilecustomtext"
	})
}
