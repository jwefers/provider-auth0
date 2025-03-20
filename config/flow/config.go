package flow

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_flow in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_flow", func(r *config.Resource) {
		r.ShortGroup = "flow"
	})
}
