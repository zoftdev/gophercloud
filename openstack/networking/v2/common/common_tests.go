package common

import (
	"gitlab.com/nxcp/tools/gophercloud"
	"gitlab.com/nxcp/tools/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
