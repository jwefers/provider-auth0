package form

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_form in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_form", func(r *config.Resource) {
		r.ShortGroup = "form"
	})
}
