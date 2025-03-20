package rule

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_rule in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_rule", func(r *config.Resource) {
		r.ShortGroup = "rule"
	})
}
