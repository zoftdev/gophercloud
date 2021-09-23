package v3

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/openstack/blockstorage/v3/qos"
	"gitlab.com/nxcp/tools/gophercloud/pagination"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestQoS(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.RequireAdmin(t)

	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	qos1, err := CreateQoS(t, client)
	th.AssertNoErr(t, err)
	defer DeleteQoS(t, client, qos1)

	qos2, err := CreateQoS(t, client)
	th.AssertNoErr(t, err)
	defer DeleteQoS(t, client, qos2)

	listOpts := qos.ListOpts{
		Limit: 1,
	}

	err = qos.List(client, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		actual, err := qos.ExtractQoS(page)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, 1, len(actual))

		var found bool
		for _, q := range actual {
			if q.ID == qos1.ID || q.ID == qos2.ID {
				found = true
			}
		}

		th.AssertEquals(t, found, true)

		return true, nil
	})

	th.AssertNoErr(t, err)

}
