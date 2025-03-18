package rule

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_rule", func(r *config.Resource) {
		r.ShortGroup = "rule"
	})
}
