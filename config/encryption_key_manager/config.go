package encryptionkeymanager

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_encryption_key_manager resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_encryption_key_manager", func(r *config.Resource) {
		r.ShortGroup = "encryptionkeymanager"
	})
}
