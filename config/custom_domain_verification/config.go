package customdomainverification

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_custom_domain_verification", func(r *config.Resource) {
		r.ShortGroup = "customdomainverification"
	})
}
