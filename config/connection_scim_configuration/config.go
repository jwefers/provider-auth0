package connectionscimconfiguration

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_scim_configuration", func(r *config.Resource) {
		r.ShortGroup = "connectionscimconfiguration"
	})
}
