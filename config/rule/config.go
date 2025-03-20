package rule

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_rule resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_rule", func(r *config.Resource) {
		r.ShortGroup = "rule"
	})
}
