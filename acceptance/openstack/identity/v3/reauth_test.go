// +build acceptance

package v3

import (
	"testing"

	"github.com/zoftdev/gophercloud/acceptance/clients"
	"github.com/zoftdev/gophercloud/openstack"
	th "github.com/zoftdev/gophercloud/testhelper"

	"github.com/zoftdev/gophercloud"
	"github.com/zoftdev/gophercloud/openstack/identity/v3/projects"
)

func TestReauthAuthResultDeadlock(t *testing.T) {
	clients.RequireAdmin(t)

	ao, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	ao.AllowReauth = true

	provider, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	provider.SetToken("this is not a valid token")

	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	pages, err := projects.List(client, nil).AllPages()
	th.AssertNoErr(t, err)
	_, err = projects.ExtractProjects(pages)
	th.AssertNoErr(t, err)
}
