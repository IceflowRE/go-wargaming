package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strconv"
)

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
//     "page" - Page 
//     "popup" - Popup window 
//     "touch" - Mobile view
// expiresAt:
//     Access_token expiration time in UNIX. Delta can also be specified in seconds.
//     Expiration time and delta must not exceed two weeks from the current time.
// language:
//     Authentication form localization language. Default is "en". Valid values:
//     
//     "cs" - Czech 
//     "de" - German 
//     "en" - English (by default)
//     "es" - Spanish 
//     "es-ar" - Argentinian Spanish 
//     "fr" - French 
//     "ja" - Japanese 
//     "pl" - Polish 
//     "pt-br" - Brazilian Portuguese 
//     "ru" - Russian 
//     "th" - Thai 
//     "tr" - Turkish 
//     "vi" - Vietnamese 
//     "zh-tw" - Traditional Chinese 
//     "zh-cn" - Simplified Chinese
// nofollow:
//     If parameter nofollow=1 is passed in, the user is not redirected. URL is returned in response. Default is 0. Min value is 0. Maximum value: 1.
// redirectUri:
//     URL where user is redirected after authentication.
//     By default: api-console.worldoftanks.com/wotx//blank/
func (client *Client) WotxAuthLogin(realm Realm, display string, expiresAt int, language string, nofollow int, redirectUri string) (*wotx.AuthLogin, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"display": display,
		"expires_at": strconv.Itoa(expiresAt),
		"language": language,
		"nofollow": strconv.Itoa(nofollow),
		"redirect_uri": redirectUri,
	}

	var result *wotx.AuthLogin
	err := client.doGetDataRequest(realm, "/wotx/auth/login/", reqParam, &result)
	return result, err
}
