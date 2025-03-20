package promptscreenrenderer

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_prompt_screen_renderer in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_screen_renderer", func(r *config.Resource) {
		r.ShortGroup = "promptscreenrenderer"
	})
}
