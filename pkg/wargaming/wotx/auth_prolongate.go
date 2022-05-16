package wotx

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type AuthProlongate struct {
	// Access_token expiration time
	ExpiresAt wgnTime.UnixTime `json:"expires_at,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Access token is passed to all methods requiring authorization.
	AccessToken string `json:"access_token,omitempty"`
} 
