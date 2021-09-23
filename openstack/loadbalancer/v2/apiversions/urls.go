package apiversions

import (
	"strings"

	"gitlab.com/nxcp/tools/gophercloud"
	"gitlab.com/nxcp/tools/gophercloud/openstack/utils"
)

func listURL(c *gophercloud.ServiceClient) string {
	baseEndpoint, _ := utils.BaseEndpoint(c.Endpoint)
	endpoint := strings.TrimRight(baseEndpoint, "/") + "/"
	return endpoint
}
