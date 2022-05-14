package wargaming

import (
	"strconv"
)

type WotAuthLogin struct {
	// URL where user is redirected for authentication.
	// This URL is returned only if parameter nofollow=1 is passed in.
	Location string `json:"location,omitempty"`
}

// WotAuthLogin Method authenticates user based on Wargaming.net ID (OpenID) which is used in World of Tanks, World of Tanks Blitz, World of Warships, World of Warplanes, and WarGag.ru. To log in, player must enter email and password used for creating account, or use a social network profile.
// Authentication is not available for iOS Game Center users in the following cases:
//   the account is not linked to a social network account, or
//   email and password are not specified in the profile.
// Information on authorization status is sent to URL specified in redirect_uri parameter.
// If authentication is successful, the following parameters are sent to redirect_uri:
// 
// status: ok — successful authentication
// access_token — access token is passed in to all methods that require authentication
// expires_at — expiration date of access_token
// account_id — user ID
// nickname — user name.
// 
// If authentication fails, the following parameters are sent to redirect_uri:
// 
// status: error — authentication error
// code — error code
// message — error message.
//
// https://developers.wargaming.net/reference/all/wot/auth/login
//
// display:
//     Layout for mobile applications. Valid values:
//     
//     "page" &mdash; Page 
//     "popup" &mdash; Popup window
// expires_at:
//     Access_token expiration time in UNIX. Delta can also be specified in seconds.
//     Expiration time and delta must not exceed two weeks from the current time.
// nofollow:
//     If parameter nofollow=1 is passed in, the user is not redirected. URL is returned in response. Default is 0. Min value is 0. Maximum value: 1.
// redirect_uri:
//     URL where user is redirected after authentication.
//     By default: api.worldoftanks.ru/wot//blank/
func (client *Client) WotAuthLogin(realm Realm, display string, expiresAt int, nofollow int, redirectUri string) (*WotAuthLogin, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"display": display,
		"expires_at": strconv.Itoa(expiresAt),
		"nofollow": strconv.Itoa(nofollow),
		"redirect_uri": redirectUri,
	}

	var result *WotAuthLogin
	err := client.doGetDataRequest(realm, "/wot/auth/login/", reqParam, &result)
	return result, err
}
