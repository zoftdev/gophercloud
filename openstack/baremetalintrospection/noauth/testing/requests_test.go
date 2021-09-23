package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/openstack/baremetalintrospection/noauth"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestNoAuth(t *testing.T) {
	noauthClient, err := noauth.NewBareMetalIntrospectionNoAuth(noauth.EndpointOpts{
		IronicInspectorEndpoint: "http://ironic:5050/v1",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", noauthClient.TokenID)
}
