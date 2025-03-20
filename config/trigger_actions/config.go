package triggeractions

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_trigger_actions in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_trigger_actions", func(r *config.Resource) {
		r.ShortGroup = "triggeractions"
	})
}
