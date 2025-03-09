// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	client "github.com/jwefers/provider-auth0/internal/controller/client/client"
	providerconfig "github.com/jwefers/provider-auth0/internal/controller/providerconfig"
	tenant "github.com/jwefers/provider-auth0/internal/controller/tenant/tenant"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		client.Setup,
		providerconfig.Setup,
		tenant.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
