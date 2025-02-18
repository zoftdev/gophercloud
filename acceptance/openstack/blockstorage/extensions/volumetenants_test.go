// +build acceptance blockstorage

package extensions

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	blockstorage "gitlab.com/nxcp/tools/gophercloud/acceptance/openstack/blockstorage/v3"
	"gitlab.com/nxcp/tools/gophercloud/openstack/blockstorage/extensions/volumetenants"
	"gitlab.com/nxcp/tools/gophercloud/openstack/blockstorage/v3/volumes"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestVolumeTenants(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")

	type volumeWithTenant struct {
		volumes.Volume
		volumetenants.VolumeTenantExt
	}

	var allVolumes []volumeWithTenant

	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	listOpts := volumes.ListOpts{
		Name: "I SHOULD NOT EXIST",
	}
	allPages, err := volumes.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	err = volumes.ExtractVolumesInto(allPages, &allVolumes)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(allVolumes))

	volume1, err := blockstorage.CreateVolume(t, client)
	th.AssertNoErr(t, err)
	defer blockstorage.DeleteVolume(t, client, volume1)

	allPages, err = volumes.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	err = volumes.ExtractVolumesInto(allPages, &allVolumes)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(allVolumes) > 0)
}
