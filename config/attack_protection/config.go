package attackprotection

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_attack_protection resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_attack_protection", func(r *config.Resource) {
		r.ShortGroup = "attackprotection"
	})
}
