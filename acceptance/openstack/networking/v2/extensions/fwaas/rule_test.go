// +build acceptance networking fwaas

package fwaas

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	"gitlab.com/nxcp/tools/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestRuleCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	rule, err := CreateRule(t, client)
	th.AssertNoErr(t, err)
	defer DeleteRule(t, client, rule.ID)

	tools.PrintResource(t, rule)

	ruleDescription := "Some rule description"
	updateOpts := rules.UpdateOpts{
		Description: &ruleDescription,
	}

	_, err = rules.Update(client, rule.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	newRule, err := rules.Get(client, rule.ID).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, newRule)

	allPages, err := rules.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allRules, err := rules.ExtractRules(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, rule := range allRules {
		if rule.ID == newRule.ID {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}
