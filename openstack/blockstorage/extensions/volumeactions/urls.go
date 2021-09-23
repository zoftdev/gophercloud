package volumeactions

import "gitlab.com/nxcp/tools/gophercloud"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
