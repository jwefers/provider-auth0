package promptscreenpartials

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_prompt_screen_partials in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_screen_partials", func(r *config.Resource) {
		r.ShortGroup = "promptscreenpartials"
	})
}
