package migrate

import (
	"github.com/zoftdev/gophercloud"
)

// MigrateResult is the response from a Migrate operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type MigrateResult struct {
	gophercloud.ErrResult
}
