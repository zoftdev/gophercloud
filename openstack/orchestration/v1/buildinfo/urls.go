package buildinfo

import "gitlab.com/nxcp/tools/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
