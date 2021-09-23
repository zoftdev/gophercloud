// +build acceptance compute availabilityzones

package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/compute/v2/extensions/availabilityzones"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestAvailabilityZonesList(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	allPages, err := availabilityzones.List(client).AllPages()
	th.AssertNoErr(t, err)

	availabilityZoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, zoneInfo := range availabilityZoneInfo {
		tools.PrintResource(t, zoneInfo)

		if zoneInfo.ZoneName == "nova" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestAvailabilityZonesListDetail(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	allPages, err := availabilityzones.ListDetail(client).AllPages()
	th.AssertNoErr(t, err)

	availabilityZoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, zoneInfo := range availabilityZoneInfo {
		tools.PrintResource(t, zoneInfo)

		if zoneInfo.ZoneName == "nova" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}
