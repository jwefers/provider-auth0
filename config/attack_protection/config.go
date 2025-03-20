package attackprotection

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_attack_protection in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_attack_protection", func(r *config.Resource) {
		r.ShortGroup = "attackprotection"
	})
}
