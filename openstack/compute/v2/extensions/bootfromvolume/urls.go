package bootfromvolume

import "github.com/zoftdev/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
