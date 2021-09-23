package serviceassets

import "github.com/zoftdev/gophercloud"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
