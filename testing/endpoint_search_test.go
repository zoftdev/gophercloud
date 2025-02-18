package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestApplyDefaultsToEndpointOpts(t *testing.T) {
	eo := gophercloud.EndpointOpts{Availability: gophercloud.AvailabilityPublic}
	eo.ApplyDefaults("compute")
	expected := gophercloud.EndpointOpts{Availability: gophercloud.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)

	eo = gophercloud.EndpointOpts{Type: "compute"}
	eo.ApplyDefaults("object-store")
	expected = gophercloud.EndpointOpts{Availability: gophercloud.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)
}
