package testing

import (
	"gitlab.com/nxcp/tools/gophercloud"
	"gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
