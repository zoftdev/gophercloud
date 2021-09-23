package limits

import (
	"gitlab.com/nxcp/tools/gophercloud"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
