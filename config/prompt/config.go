package prompt

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_prompt in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt", func(r *config.Resource) {
		r.ShortGroup = "prompt"
	})
}
