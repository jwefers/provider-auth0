package promptscreenpartial

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_prompt_screen_partial resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_screen_partial", func(r *config.Resource) {
		r.ShortGroup = "promptscreenpartial"
	})
}
