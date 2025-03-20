package promptscreenpartial

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_prompt_screen_partial in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_screen_partial", func(r *config.Resource) {
		r.ShortGroup = "promptscreenpartial"
	})
}
