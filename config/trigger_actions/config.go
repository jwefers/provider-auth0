package triggeractions

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_trigger_actions resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_trigger_actions", func(r *config.Resource) {
		r.ShortGroup = "triggeractions"
	})
}
