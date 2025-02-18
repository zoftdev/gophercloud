// +build acceptance keymanager orders

package v1

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/keymanager/v1/containers"
	"gitlab.com/nxcp/tools/gophercloud/openstack/keymanager/v1/orders"
	"gitlab.com/nxcp/tools/gophercloud/openstack/keymanager/v1/secrets"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestOrdersCRUD(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/queens")
	clients.RequireAdmin(t)

	client, err := clients.NewKeyManagerV1Client()
	th.AssertNoErr(t, err)

	order, err := CreateKeyOrder(t, client)
	th.AssertNoErr(t, err)
	orderID, err := ParseID(order.OrderRef)
	th.AssertNoErr(t, err)
	defer DeleteOrder(t, client, orderID)

	secretID, err := ParseID(order.SecretRef)
	th.AssertNoErr(t, err)

	payloadOpts := secrets.GetPayloadOpts{
		PayloadContentType: "application/octet-stream",
	}
	payload, err := secrets.GetPayload(client, secretID, payloadOpts).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, payload)

	allPages, err := orders.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allOrders, err := orders.ExtractOrders(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, v := range allOrders {
		if v.OrderRef == order.OrderRef {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestOrdersAsymmetric(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/queens")
	clients.RequireAdmin(t)

	client, err := clients.NewKeyManagerV1Client()
	th.AssertNoErr(t, err)

	order, err := CreateAsymmetricOrder(t, client)
	th.AssertNoErr(t, err)
	orderID, err := ParseID(order.OrderRef)
	th.AssertNoErr(t, err)
	defer DeleteOrder(t, client, orderID)

	containerID, err := ParseID(order.ContainerRef)
	th.AssertNoErr(t, err)

	container, err := containers.Get(client, containerID).Extract()
	th.AssertNoErr(t, err)

	for _, v := range container.SecretRefs {
		secretID, err := ParseID(v.SecretRef)
		th.AssertNoErr(t, err)

		payloadOpts := secrets.GetPayloadOpts{
			PayloadContentType: "application/octet-stream",
		}

		payload, err := secrets.GetPayload(client, secretID, payloadOpts).Extract()
		th.AssertNoErr(t, err)
		tools.PrintResource(t, string(payload))
	}
}
