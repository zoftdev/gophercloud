// +build acceptance networking provider

package extensions

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	networking "gitlab.com/nxcp/tools/gophercloud/acceptance/openstack/networking/v2"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/networking/v2/networks"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestNetworksProviderCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create a network
	network, err := networking.CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer networking.DeleteNetwork(t, client, network.ID)

	getResult := networks.Get(client, network.ID)
	newNetwork, err := getResult.Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newNetwork)
}
