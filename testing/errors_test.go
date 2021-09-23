package testing

import (
	"testing"

	"gitlab.com/nxcp/tools/gophercloud"
	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
)

func TestGetResponseCode(t *testing.T) {
	respErr := gophercloud.ErrUnexpectedResponseCode{
		URL:            "http://example.com",
		Method:         "GET",
		Expected:       []int{200},
		Actual:         404,
		Body:           nil,
		ResponseHeader: nil,
	}

	var err404 error = gophercloud.ErrDefault404{ErrUnexpectedResponseCode: respErr}

	err, ok := err404.(gophercloud.StatusCodeError)
	th.AssertEquals(t, true, ok)
	th.AssertEquals(t, err.GetStatusCode(), 404)
}
