package testing

import (
	"net/http"
	"testing"

	th "gitlab.com/nxcp/tools/gophercloud/testhelper"
	"gitlab.com/nxcp/tools/gophercloud/testhelper/client"
)

func mockPauseServerResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"pause": null}`)
		w.WriteHeader(http.StatusAccepted)
	})
}

func mockUnpauseServerResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"unpause": null}`)
		w.WriteHeader(http.StatusAccepted)
	})
}
