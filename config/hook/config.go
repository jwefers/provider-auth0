package hook

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_hook resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_hook", func(r *config.Resource) {
		r.ShortGroup = "hook"
	})
}
