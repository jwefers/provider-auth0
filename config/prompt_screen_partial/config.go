package promptscreenpartial

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_screen_partial", func(r *config.Resource) {
		r.ShortGroup = "promptscreenpartial"
	})
}
