package action

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_action in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_action", func(r *config.Resource) {
		r.ShortGroup = "action"
	})
}
