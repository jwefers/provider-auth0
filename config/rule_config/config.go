package ruleconfig

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_rule_config in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_rule_config", func(r *config.Resource) {
		r.ShortGroup = "ruleconfig"
	})
}
