// +build acceptance containers capsules

package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/loadbalancer/v2/amphorae"
)

func TestAmphoraeList(t *testing.T) {
	clients.RequireAdmin(t)
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")
	clients.SkipRelease(t, "stable/pike")
	clients.SkipRelease(t, "stable/queens")
	clients.SkipRelease(t, "stable/rocky")

	client, err := clients.NewLoadBalancerV2Client()
	if err != nil {
		t.Fatalf("Unable to create a loadbalancer client: %v", err)
	}

	allPages, err := amphorae.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list amphorae: %v", err)
	}

	allAmphorae, err := amphorae.ExtractAmphorae(allPages)
	if err != nil {
		t.Fatalf("Unable to extract amphorae: %v", err)
	}

	for _, amphora := range allAmphorae {
		tools.PrintResource(t, amphora)
	}
}
