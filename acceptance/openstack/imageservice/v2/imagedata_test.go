package v2

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/acceptance/clients"
	"gitlab.com/nxcp/tools/gophercloud/acceptance/tools"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestImageStage(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")
	clients.SkipRelease(t, "stable/pike")
	clients.SkipRelease(t, "stable/queens")
	clients.SkipRelease(t, "stable/rocky")

	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	image, err := CreateEmptyImage(t, client)
	th.AssertNoErr(t, err)
	defer DeleteImage(t, client, image)

	imageFileName := tools.RandomString("image_", 8)
	imageFilepath := "/tmp/" + imageFileName
	imageURL := ImportImageURL

	err = DownloadImageFileFromURL(t, imageURL, imageFilepath)
	th.AssertNoErr(t, err)
	defer DeleteImageFile(t, imageFilepath)

	err = StageImage(t, client, imageFilepath, image.ID)
	th.AssertNoErr(t, err)
}
