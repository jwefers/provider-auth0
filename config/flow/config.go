package flow

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_flow resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_flow", func(r *config.Resource) {
		r.ShortGroup = "flow"
	})
}
