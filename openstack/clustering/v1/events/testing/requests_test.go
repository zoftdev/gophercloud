package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/openstack/clustering/v1/events"

	"gitlab.com/nxcp/tools/gophercloud/pagination"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
	fake "gitlab.com/nxcp/tools/gophercloud/testhelper/client"
)

func TestListEvents(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListSuccessfully(t)

	pageCount := 0
	err := events.List(fake.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		pageCount++
		actual, err := events.ExtractEvents(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedEvents, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)

	if pageCount != 1 {
		t.Errorf("Expected 1 page, got %d", pageCount)
	}
}

func TestGetEvent(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t, ExpectedEvent1.ID)

	actual, err := events.Get(fake.ServiceClient(), ExpectedEvent1.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedEvent1, *actual)
}
