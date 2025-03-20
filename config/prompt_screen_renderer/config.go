package promptscreenrenderer

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_prompt_screen_renderer resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_screen_renderer", func(r *config.Resource) {
		r.ShortGroup = "promptscreenrenderer"
	})
}
