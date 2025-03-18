package selfserviceprofile

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_self_service_profile", func(r *config.Resource) {
		r.ShortGroup = "selfserviceprofile"
	})
}
