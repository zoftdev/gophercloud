package webhooks

import "github.com/zoftdev/gophercloud"

func triggerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("v1", "webhooks", id, "trigger")
}
