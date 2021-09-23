package extensions

import "gitlab.com/nxcp/tools/gophercloud"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
