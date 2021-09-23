package services

import "gitlab.com/nxcp/tools/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
