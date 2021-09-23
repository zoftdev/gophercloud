package swauth

import "gitlab.com/nxcp/tools/gophercloud"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
