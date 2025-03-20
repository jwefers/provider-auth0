package promptpartials

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_prompt_partials resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_partials", func(r *config.Resource) {
		r.ShortGroup = "promptpartials"
	})
}
