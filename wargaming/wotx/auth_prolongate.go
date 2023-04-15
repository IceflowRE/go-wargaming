// Auto generated file!

package wotx

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AuthProlongateOptions struct {
	// Access_token expiration time in UNIX. Delta can also be specified in seconds.
	// Expiration time and delta must not exceed two weeks from the current time.
	ExpiresAt *wgnTime.UnixTime `json:"expires_at,omitempty"`
}

type AuthProlongate struct {
	// Access token is passed to all methods requiring authorization.
	AccessToken *string `json:"access_token,omitempty"`
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Access_token expiration time
	ExpiresAt *wgnTime.UnixTime `json:"expires_at,omitempty"`
}
