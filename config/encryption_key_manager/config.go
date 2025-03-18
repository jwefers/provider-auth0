package encryptionkeymanager

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_encryption_key_manager", func(r *config.Resource) {
		r.ShortGroup = "encryptionkeymanager"
	})
}
