// +build acceptance imageservice imageimport

package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestGetImportInfo(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")
	clients.SkipRelease(t, "stable/pike")
	clients.SkipRelease(t, "stable/queens")

	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	importInfo, err := GetImportInfo(t, client)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, importInfo)
}

func TestCreateImport(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")
	clients.SkipRelease(t, "stable/pike")
	clients.SkipRelease(t, "stable/queens")

	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	image, err := CreateEmptyImage(t, client)
	th.AssertNoErr(t, err)
	defer DeleteImage(t, client, image)

	err = ImportImage(t, client, image.ID)
	th.AssertNoErr(t, err)
}
