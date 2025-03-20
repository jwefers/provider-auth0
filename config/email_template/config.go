package emailtemplate

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_email_template resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_email_template", func(r *config.Resource) {
		r.ShortGroup = "emailtemplate"
	})
}
