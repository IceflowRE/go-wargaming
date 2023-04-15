// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strconv"
)

// AuthLogin authenticates user based on PlayStation Network ID which is used in World of Tanks PlayStation 4. To log in, player must enter PlayStation Network ID and password.
// Information on authorization status is sent to URL specified in redirect_uri parameter.
// If authentication is successful, the following parameters are sent to redirect_uri:
//
// status: ok - successful authentication
// access_token - access token is passed in to all methods that require authentication
// expires_at - expiration date of access_token
// account_id - user ID
// nickname - user name
// psn_access_token - PlayStation Network access token
// psn_access_token_expires_at - expiration date of psn_access_token.
//
// If authentication fails, the following parameters are sent to redirect_uri:
//
// status: error - authentication error
// code - error code
// message - error message.
//
// https://developers.wargaming.net/reference/all/wotx/auth/login
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) AuthLogin(ctx context.Context, realm Realm, options *wotx.AuthLoginOptions) (*wotx.AuthLogin, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Display != nil {
			reqParam["display"] = *options.Display
		}
		if options.ExpiresAt != nil {
			reqParam["expires_at"] = strconv.FormatInt(options.ExpiresAt.Unix(), 10)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Nofollow != nil {
			reqParam["nofollow"] = strconv.Itoa(*options.Nofollow)
		}
		if options.RedirectUri != nil {
			reqParam["redirect_uri"] = *options.RedirectUri
		}
	}

	var data *wotx.AuthLogin
	err := service.client.getRequest(ctx, sectionWotx, realm, "/auth/login/", reqParam, &data, nil)
	return data, err
}
