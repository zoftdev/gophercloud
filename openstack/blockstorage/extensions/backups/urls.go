package backups

import "gitlab.com/nxcp/tools/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("backups")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("backups")
}

func listDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("backups", "detail")
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func restoreURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("backups", id, "restore")
}

func exportURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("backups", id, "export_record")
}

func importURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("backups", "import_record")
}
