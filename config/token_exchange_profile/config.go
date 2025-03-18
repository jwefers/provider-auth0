package tokenexchangeprofile

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_token_exchange_profile", func(r *config.Resource) {
		r.ShortGroup = "tokenexchangeprofile"
	})
}
