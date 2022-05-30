package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strconv"
)

// WotxAuthProlongate Method generates new access_token based on the current token.
// This method is used when the player is still using the application but the current access_token is about to expire.
//
// https://developers.wargaming.net/reference/all/wotx/auth/prolongate
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//     Parameter is required.
// expiresAt:
//     Access_token expiration time in UNIX. Delta can also be specified in seconds.
//     Expiration time and delta must not exceed two weeks from the current time.
func (client *Client) WotxAuthProlongate(realm Realm, accessToken string, expiresAt int) (*wotx.AuthProlongate, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"expires_at": strconv.Itoa(expiresAt),
	}

	var result *wotx.AuthProlongate
	err := client.doPostDataRequest(realm, "/wotx/auth/prolongate/", reqParam, &result)
	return result, err
}
