// +build acceptance identity

package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack"
	"gitlab.com/nxcp/tools/gophercloud/openstack/identity/v2/tokens"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestTokenAuthenticate(t *testing.T) {
	clients.RequireIdentityV2(t)
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV2UnauthenticatedClient()
	th.AssertNoErr(t, err)

	authOptions, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	result := tokens.Create(client, authOptions)
	token, err := result.ExtractToken()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, token)

	catalog, err := result.ExtractServiceCatalog()
	th.AssertNoErr(t, err)

	for _, entry := range catalog.Entries {
		tools.PrintResource(t, entry)
	}
}

func TestTokenValidate(t *testing.T) {
	clients.RequireIdentityV2(t)
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV2Client()
	th.AssertNoErr(t, err)

	authOptions, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	result := tokens.Create(client, authOptions)
	token, err := result.ExtractToken()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, token)

	getResult := tokens.Get(client, token.ID)
	user, err := getResult.ExtractUser()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, user)
}
