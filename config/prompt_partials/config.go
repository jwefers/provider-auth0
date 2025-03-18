package promptpartials

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_prompt_partials", func(r *config.Resource) {
		r.ShortGroup = "promptpartials"
	})
}
