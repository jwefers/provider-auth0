package rolepermissions

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_role_permissions", func(r *config.Resource) {
		r.ShortGroup = "rolepermissions"
		/* TODO: below asigment of empty array is a workaround for generating the example manifest for this resource
		Following error is thrown when running 'make generate' and needs to be investigated:
			panic: cannot store examples: cannot store example manifest for resource: auth0_role_permissions.*:
			cannot resolve references of resource: failed to resolve references of paved:
			cannot get string value from paved: spec.forProvider.scopes: spec.forProvider.scopes: not a string
		*/
		r.MetaResource.Examples = []registry.ResourceExample{}
	})
}
