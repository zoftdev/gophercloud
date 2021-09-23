package availabilityzones

import "gitlab.com/nxcp/tools/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
