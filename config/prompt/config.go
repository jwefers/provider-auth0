package prompt

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_prompt resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt", func(r *config.Resource) {
		r.ShortGroup = "prompt"
	})
}
