package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/openstack/blockstorage/extensions/services"
	"gitlab.com/nxcp/tools/gophercloud/pagination"
	"gitlab.com/nxcp/tools/gophercloud/testhelper"
	"gitlab.com/nxcp/tools/gophercloud/testhelper/client"
)

func TestListServices(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()
	HandleListSuccessfully(t)

	pages := 0
	err := services.List(client.ServiceClient(), services.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := services.ExtractServices(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 services, got %d", len(actual))
		}
		testhelper.CheckDeepEquals(t, FirstFakeService, actual[0])
		testhelper.CheckDeepEquals(t, SecondFakeService, actual[1])
		testhelper.CheckDeepEquals(t, ThirdFakeService, actual[2])

		return true, nil
	})

	testhelper.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}
