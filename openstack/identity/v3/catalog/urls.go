package catalog

import "gitlab.com/nxcp/tools/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}
