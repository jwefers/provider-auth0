package emailtemplate

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_email_template in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_email_template", func(r *config.Resource) {
		r.ShortGroup = "emailtemplate"
	})
}
