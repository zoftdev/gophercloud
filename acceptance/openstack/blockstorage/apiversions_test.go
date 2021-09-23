// +build acceptance blockstorage

package blockstorage

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/blockstorage/apiversions"
)

func TestAPIVersionsList(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	allPages, err := apiversions.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve API versions: %v", err)
	}

	allVersions, err := apiversions.ExtractAPIVersions(allPages)
	if err != nil {
		t.Fatalf("Unable to extract API versions: %v", err)
	}

	for _, v := range allVersions {
		tools.PrintResource(t, v)
	}
}

func TestAPIVersionsGet(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	allPages, err := apiversions.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve API versions: %v", err)
	}

	v, err := apiversions.ExtractAPIVersion(allPages, "v3.0")
	if err != nil {
		t.Fatalf("Unable to extract API version: %v", err)
	}

	tools.PrintResource(t, v)
}
