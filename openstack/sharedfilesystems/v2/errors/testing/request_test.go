package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud/openstack/sharedfilesystems/v2/errors"
	"gitlab.com/nxcp/tools/gophercloud/openstack/sharedfilesystems/v2/shares"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
	"gitlab.com/nxcp/tools/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &shares.CreateOpts{Size: 1, Name: "my_test_share", ShareProto: "NFS", SnapshotID: "70bfbebc-d3ff-4528-8bbb-58422daa280b"}
	_, err := shares.Create(client.ServiceClient(), options).Extract()

	if err == nil {
		t.Fatal("Expected error")
	}

	detailedErr := errors.ErrorDetails{}
	e := errors.ExtractErrorInto(err, &detailedErr)
	th.AssertNoErr(t, e)

	for k, msg := range detailedErr {
		th.AssertEquals(t, k, "itemNotFound")
		th.AssertEquals(t, msg.Code, 404)
		th.AssertEquals(t, msg.Message, "ShareSnapshotNotFound: Snapshot 70bfbebc-d3ff-4528-8bbb-58422daa280b could not be found.")
	}
}
