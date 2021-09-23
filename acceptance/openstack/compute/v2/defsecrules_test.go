// +build acceptance compute defsecrules

package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	dsr "gitlab.com/nxcp/tools/gophercloud/openstack/compute/v2/extensions/defsecrules"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestDefSecRulesList(t *testing.T) {
	clients.RequireAdmin(t)
	clients.RequireNovaNetwork(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	allPages, err := dsr.List(client).AllPages()
	th.AssertNoErr(t, err)

	allDefaultRules, err := dsr.ExtractDefaultRules(allPages)
	th.AssertNoErr(t, err)

	for _, defaultRule := range allDefaultRules {
		tools.PrintResource(t, defaultRule)
	}
}

func TestDefSecRulesCreate(t *testing.T) {
	clients.RequireAdmin(t)
	clients.RequireNovaNetwork(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	defaultRule, err := CreateDefaultRule(t, client)
	th.AssertNoErr(t, err)
	defer DeleteDefaultRule(t, client, defaultRule)

	tools.PrintResource(t, defaultRule)
}

func TestDefSecRulesGet(t *testing.T) {
	clients.RequireAdmin(t)
	clients.RequireNovaNetwork(t)

	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	defaultRule, err := CreateDefaultRule(t, client)
	th.AssertNoErr(t, err)
	defer DeleteDefaultRule(t, client, defaultRule)

	newDefaultRule, err := dsr.Get(client, defaultRule.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newDefaultRule)
}
