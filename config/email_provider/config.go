package emailprovider

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_email_provider", func(r *config.Resource) {
		r.ShortGroup = "emailprovider"
	})
}
