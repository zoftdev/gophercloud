// +build acceptance blockstorage

package noauth

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/blockstorage/v2/volumes"
)

func TestVolumesList(t *testing.T) {
	client, err := clients.NewBlockStorageV2NoAuthClient()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	allPages, err := volumes.List(client, volumes.ListOpts{}).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve volumes: %v", err)
	}

	allVolumes, err := volumes.ExtractVolumes(allPages)
	if err != nil {
		t.Fatalf("Unable to extract volumes: %v", err)
	}

	for _, volume := range allVolumes {
		tools.PrintResource(t, volume)
	}
}

func TestVolumesCreateDestroy(t *testing.T) {
	client, err := clients.NewBlockStorageV2NoAuthClient()
	if err != nil {
		t.Fatalf("Unable to create blockstorage client: %v", err)
	}

	volume, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}
	defer DeleteVolume(t, client, volume)

	newVolume, err := volumes.Get(client, volume.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve volume: %v", err)
	}

	tools.PrintResource(t, newVolume)
}
