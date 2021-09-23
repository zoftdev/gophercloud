package bootfromvolume

import "gitlab.com/nxcp/tools/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
