package logstream

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_log_stream", func(r *config.Resource) {
		r.ShortGroup = "logstream"
	})
}
