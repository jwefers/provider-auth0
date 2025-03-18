package connectionclient

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_client", func(r *config.Resource) {
		r.ShortGroup = "connectionclient"
	})
}
