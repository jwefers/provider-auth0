package hook

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_hook in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_hook", func(r *config.Resource) {
		r.ShortGroup = "hook"
	})
}
