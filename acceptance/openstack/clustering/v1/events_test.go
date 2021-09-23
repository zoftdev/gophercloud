// +build acceptance clustering events

package v1

import (
	"testing"

	"github.com/zoftdev/gophercloud/acceptance/clients"
	"github.com/zoftdev/gophercloud/acceptance/tools"
	"github.com/zoftdev/gophercloud/openstack/clustering/v1/events"
	th "github.com/zoftdev/gophercloud/testhelper"
)

func TestEventsList(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	opts := events.ListOpts{
		Limit: 200,
	}

	allPages, err := events.List(client, opts).AllPages()
	th.AssertNoErr(t, err)

	allEvents, err := events.ExtractEvents(allPages)
	th.AssertNoErr(t, err)

	for _, event := range allEvents {
		tools.PrintResource(t, event)
	}
}
