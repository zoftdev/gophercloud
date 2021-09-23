package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/openstack/baremetal/noauth"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestNoAuth(t *testing.T) {
	noauthClient, err := noauth.NewBareMetalNoAuth(noauth.EndpointOpts{
		IronicEndpoint: "http://ironic:6385/v1",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", noauthClient.TokenID)
}
