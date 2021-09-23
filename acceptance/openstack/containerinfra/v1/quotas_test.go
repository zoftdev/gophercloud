// +build acceptance containerinfra

package v1

import (
	"testing"

	"github.com/zoftdev/gophercloud/acceptance/clients"
	"github.com/zoftdev/gophercloud/acceptance/tools"
	th "github.com/zoftdev/gophercloud/testhelper"
)

func TestQuotasCRUD(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	th.AssertNoErr(t, err)

	quota, err := CreateQuota(t, client)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, quota)
}
