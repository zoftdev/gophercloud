package quotas

import (
	"gitlab.com/nxcp/tools/gophercloud"
)

var apiName = "quotas"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}
