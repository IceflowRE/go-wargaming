package wargaming

import (
	"strconv"
)

type WotxAuthLogin struct {
	// URL where user is redirected for authentication.
	// This URL is returned only if parameter nofollow=1 is passed in.
	Location string `json:"location,omitempty"`
}

// WotxAuthLogin Method authenticates user based on PlayStation Network ID which is used in World of Tanks PlayStation 4. To log in, player must enter PlayStation Network ID and password.
// Information on authorization status is sent to URL specified in redirect_uri parameter.
// If authentication is successful, the following parameters are sent to redirect_uri:
// 
// status: ok — successful authentication
// access_token — access token is passed in to all methods that require authentication
// expires_at — expiration date of access_token
// account_id — user ID
// nickname — user name
// psn_access_token — PlayStation Network access token
// psn_access_token_expires_at — expiration date of psn_access_token.
// 
// If authentication fails, the following parameters are sent to redirect_uri:
// 
// status: error — authentication error
// code — error code
// message — error message.
//
// https://developers.wargaming.net/reference/all/wotx/auth/login
//
// display:
//     Layout for mobile applications. Valid values:
//     
//     "page" &mdash; Page 
//     "popup" &mdash; Popup window 
//     "touch" &mdash; Mobile view
// expires_at:
//     Access_token expiration time in UNIX. Delta can also be specified in seconds.
//     Expiration time and delta must not exceed two weeks from the current time.
// language:
//     Authentication form localization language. Default is "en". Valid values:
//     
//     "cs" &mdash; Czech 
//     "de" &mdash; German 
//     "en" &mdash; English (by default)
//     "es" &mdash; Spanish 
//     "es-ar" &mdash; Argentinian Spanish 
//     "fr" &mdash; French 
//     "ja" &mdash; Japanese 
//     "pl" &mdash; Polish 
//     "pt-br" &mdash; Brazilian Portuguese 
//     "ru" &mdash; Russian 
//     "th" &mdash; Thai 
//     "tr" &mdash; Turkish 
//     "vi" &mdash; Vietnamese 
//     "zh-tw" &mdash; Traditional Chinese 
//     "zh-cn" &mdash; Simplified Chinese
// nofollow:
//     If parameter nofollow=1 is passed in, the user is not redirected. URL is returned in response. Default is 0. Min value is 0. Maximum value: 1.
// redirect_uri:
//     URL where user is redirected after authentication.
//     By default: api-console.worldoftanks.com/wotx//blank/
func (client *Client) WotxAuthLogin(realm Realm, display string, expiresAt int, language string, nofollow int, redirectUri string) (*WotxAuthLogin, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"display": display,
		"expires_at": strconv.Itoa(expiresAt),
		"language": language,
		"nofollow": strconv.Itoa(nofollow),
		"redirect_uri": redirectUri,
	}

	var result *WotxAuthLogin
	err := client.doGetDataRequest(realm, "/wotx/auth/login/", reqParam, &result)
	return result, err
}
