package triggeraction

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_trigger_action resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_trigger_action", func(r *config.Resource) {
		r.ShortGroup = "triggeraction"
	})
}
