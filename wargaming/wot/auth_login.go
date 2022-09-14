// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// AuthLoginOptions options.
type AuthLoginOptions struct {
	// Layout for mobile applications. Valid values:
	//
	// "page" - Page
	// "popup" - Popup window
	Display *string
	// Access_token expiration time in UNIX. Delta can also be specified in seconds.
	// Expiration time and delta must not exceed two weeks from the current time.
	ExpiresAt *wgnTime.UnixTime
	// If parameter nofollow=1 is passed in, the user is not redirected. URL is returned in response. Default is 0. Min value is 0. Maximum value: 1.
	Nofollow *int
	// URL where user is redirected after authentication.
	// By default: api.worldoftanks.ru/wot//blank/
	RedirectUri *string
}

type AuthLogin struct {
	// URL where user is redirected for authentication.
	// This URL is returned only if parameter nofollow=1 is passed in.
	Location *string `json:"location,omitempty"`
}
