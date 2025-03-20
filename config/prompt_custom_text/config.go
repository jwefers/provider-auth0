package promptcustomtext

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_prompt_custom_text resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_custom_text", func(r *config.Resource) {
		r.ShortGroup = "promptcustomtext"
	})
}
