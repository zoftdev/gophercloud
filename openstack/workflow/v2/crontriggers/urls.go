package crontriggers

import "gitlab.com/nxcp/tools/gophercloud"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("cron_triggers")
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("cron_triggers", id)
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("cron_triggers", id)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("cron_triggers")
}
