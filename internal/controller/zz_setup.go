// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	action "github.com/jwefers/provider-auth0/internal/controller/action/action"
	protection "github.com/jwefers/provider-auth0/internal/controller/attackprotection/protection"
	branding "github.com/jwefers/provider-auth0/internal/controller/branding/branding"
	theme "github.com/jwefers/provider-auth0/internal/controller/brandingtheme/theme"
	client "github.com/jwefers/provider-auth0/internal/controller/client/client"
	credentials "github.com/jwefers/provider-auth0/internal/controller/clientcredentials/credentials"
	grant "github.com/jwefers/provider-auth0/internal/controller/clientgrant/grant"
	connection "github.com/jwefers/provider-auth0/internal/controller/connection/connection"
	clientconnectionclient "github.com/jwefers/provider-auth0/internal/controller/connectionclient/client"
	clients "github.com/jwefers/provider-auth0/internal/controller/connectionclients/clients"
	scimconfiguration "github.com/jwefers/provider-auth0/internal/controller/connectionscimconfiguration/scimconfiguration"
	domain "github.com/jwefers/provider-auth0/internal/controller/customdomain/domain"
	domainverification "github.com/jwefers/provider-auth0/internal/controller/customdomainverification/domainverification"
	provider "github.com/jwefers/provider-auth0/internal/controller/emailprovider/provider"
	template "github.com/jwefers/provider-auth0/internal/controller/emailtemplate/template"
	keymanager "github.com/jwefers/provider-auth0/internal/controller/encryptionkeymanager/keymanager"
	flow "github.com/jwefers/provider-auth0/internal/controller/flow/flow"
	vaultconnection "github.com/jwefers/provider-auth0/internal/controller/flowvaultconnection/vaultconnection"
	form "github.com/jwefers/provider-auth0/internal/controller/form/form"
	guardian "github.com/jwefers/provider-auth0/internal/controller/guardian/guardian"
	hook "github.com/jwefers/provider-auth0/internal/controller/hook/hook"
	stream "github.com/jwefers/provider-auth0/internal/controller/logstream/stream"
	organization "github.com/jwefers/provider-auth0/internal/controller/organization/organization"
	clientgrant "github.com/jwefers/provider-auth0/internal/controller/organizationclientgrant/clientgrant"
	connectionorganizationconnection "github.com/jwefers/provider-auth0/internal/controller/organizationconnection/connection"
	connections "github.com/jwefers/provider-auth0/internal/controller/organizationconnections/connections"
	member "github.com/jwefers/provider-auth0/internal/controller/organizationmember/member"
	memberrole "github.com/jwefers/provider-auth0/internal/controller/organizationmemberrole/memberrole"
	memberroles "github.com/jwefers/provider-auth0/internal/controller/organizationmemberroles/memberroles"
	members "github.com/jwefers/provider-auth0/internal/controller/organizationmembers/members"
	pages "github.com/jwefers/provider-auth0/internal/controller/pages/pages"
	providerphoneprovider "github.com/jwefers/provider-auth0/internal/controller/phoneprovider/provider"
	prompt "github.com/jwefers/provider-auth0/internal/controller/prompt/prompt"
	customtext "github.com/jwefers/provider-auth0/internal/controller/promptcustomtext/customtext"
	partials "github.com/jwefers/provider-auth0/internal/controller/promptpartials/partials"
	screenpartial "github.com/jwefers/provider-auth0/internal/controller/promptscreenpartial/screenpartial"
	screenpartials "github.com/jwefers/provider-auth0/internal/controller/promptscreenpartials/screenpartials"
	screenrenderer "github.com/jwefers/provider-auth0/internal/controller/promptscreenrenderer/screenrenderer"
	providerconfig "github.com/jwefers/provider-auth0/internal/controller/providerconfig"
	server "github.com/jwefers/provider-auth0/internal/controller/resourceserver/server"
	serverscope "github.com/jwefers/provider-auth0/internal/controller/resourceserverscope/serverscope"
	serverscopes "github.com/jwefers/provider-auth0/internal/controller/resourceserverscopes/serverscopes"
	role "github.com/jwefers/provider-auth0/internal/controller/role/role"
	permission "github.com/jwefers/provider-auth0/internal/controller/rolepermission/permission"
	permissions "github.com/jwefers/provider-auth0/internal/controller/rolepermissions/permissions"
	rule "github.com/jwefers/provider-auth0/internal/controller/rule/rule"
	config "github.com/jwefers/provider-auth0/internal/controller/ruleconfig/config"
	serviceprofile "github.com/jwefers/provider-auth0/internal/controller/selfserviceprofile/serviceprofile"
	serviceprofilecustomtext "github.com/jwefers/provider-auth0/internal/controller/selfserviceprofilecustomtext/serviceprofilecustomtext"
	tenant "github.com/jwefers/provider-auth0/internal/controller/tenant/tenant"
	exchangeprofile "github.com/jwefers/provider-auth0/internal/controller/tokenexchangeprofile/exchangeprofile"
	actiontriggeraction "github.com/jwefers/provider-auth0/internal/controller/triggeraction/action"
	actions "github.com/jwefers/provider-auth0/internal/controller/triggeractions/actions"
	user "github.com/jwefers/provider-auth0/internal/controller/user/user"
	permissionuserpermission "github.com/jwefers/provider-auth0/internal/controller/userpermission/permission"
	permissionsuserpermissions "github.com/jwefers/provider-auth0/internal/controller/userpermissions/permissions"
	roleuserrole "github.com/jwefers/provider-auth0/internal/controller/userrole/role"
	roles "github.com/jwefers/provider-auth0/internal/controller/userroles/roles"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		action.Setup,
		protection.Setup,
		branding.Setup,
		theme.Setup,
		client.Setup,
		credentials.Setup,
		grant.Setup,
		connection.Setup,
		clientconnectionclient.Setup,
		clients.Setup,
		scimconfiguration.Setup,
		domain.Setup,
		domainverification.Setup,
		provider.Setup,
		template.Setup,
		keymanager.Setup,
		flow.Setup,
		vaultconnection.Setup,
		form.Setup,
		guardian.Setup,
		hook.Setup,
		stream.Setup,
		organization.Setup,
		clientgrant.Setup,
		connectionorganizationconnection.Setup,
		connections.Setup,
		member.Setup,
		memberrole.Setup,
		memberroles.Setup,
		members.Setup,
		pages.Setup,
		providerphoneprovider.Setup,
		prompt.Setup,
		customtext.Setup,
		partials.Setup,
		screenpartial.Setup,
		screenpartials.Setup,
		screenrenderer.Setup,
		providerconfig.Setup,
		server.Setup,
		serverscope.Setup,
		serverscopes.Setup,
		role.Setup,
		permission.Setup,
		permissions.Setup,
		rule.Setup,
		config.Setup,
		serviceprofile.Setup,
		serviceprofilecustomtext.Setup,
		tenant.Setup,
		exchangeprofile.Setup,
		actiontriggeraction.Setup,
		actions.Setup,
		user.Setup,
		permissionuserpermission.Setup,
		permissionsuserpermissions.Setup,
		roleuserrole.Setup,
		roles.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
