package ec2credentials

import "gitlab.com/nxcp/tools/gophercloud"

func listURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2")
}

func getURL(client *gophercloud.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2", id)
}

func createURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2")
}

func deleteURL(client *gophercloud.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2", id)
}
