// +build acceptance compute remoteconsoles

package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestRemoteConsoleCreate(t *testing.T) {
	clients.RequireLong(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	client.Microversion = "2.6"

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	remoteConsole, err := CreateRemoteConsole(t, client, server.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, remoteConsole)
}
