package credentials

import "gitlab.com/nxcp/tools/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("credentials")
}

func getURL(client *gophercloud.ServiceClient, credentialID string) string {
	return client.ServiceURL("credentials", credentialID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("credentials")
}

func deleteURL(client *gophercloud.ServiceClient, credentialID string) string {
	return client.ServiceURL("credentials", credentialID)
}

func updateURL(client *gophercloud.ServiceClient, credentialID string) string {
	return client.ServiceURL("credentials", credentialID)
}
