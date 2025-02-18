// +build acceptance

package v1

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/orchestration/v1/stackresources"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestStackResources(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")

	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	stack, err := CreateStack(t, client)
	th.AssertNoErr(t, err)
	defer DeleteStack(t, client, stack.Name, stack.ID)

	resource, err := stackresources.Get(client, stack.Name, stack.ID, basicTemplateResourceName).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, resource)

	metadata, err := stackresources.Metadata(client, stack.Name, stack.ID, basicTemplateResourceName).Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, metadata)

	markUnhealthyOpts := &stackresources.MarkUnhealthyOpts{
		MarkUnhealthy:        true,
		ResourceStatusReason: "Wrong security policy is detected.",
	}

	err = stackresources.MarkUnhealthy(client, stack.Name, stack.ID, basicTemplateResourceName, markUnhealthyOpts).ExtractErr()
	th.AssertNoErr(t, err)

	unhealthyResource, err := stackresources.Get(client, stack.Name, stack.ID, basicTemplateResourceName).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "CHECK_FAILED", unhealthyResource.Status)
	tools.PrintResource(t, unhealthyResource)

	allPages, err := stackresources.List(client, stack.Name, stack.ID, nil).AllPages()
	th.AssertNoErr(t, err)
	allResources, err := stackresources.ExtractResources(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, v := range allResources {
		if v.Name == basicTemplateResourceName {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}
