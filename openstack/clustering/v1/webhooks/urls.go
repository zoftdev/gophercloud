package webhooks

import "gitlab.com/nxcp/tools/gophercloud"

func triggerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("v1", "webhooks", id, "trigger")
}
