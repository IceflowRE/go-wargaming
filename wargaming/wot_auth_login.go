// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
)

// AuthLogin authenticates user based on Wargaming.net ID (OpenID) which is used in World of Tanks, World of Tanks Blitz, World of Warships, World of Warplanes, and WarGag.ru. To log in, player must enter email and password used for creating account, or use a social network profile.
// Authentication is not available for iOS Game Center users in the following cases:
//   the account is not linked to a social network account, or
//   email and password are not specified in the profile.
// Information on authorization status is sent to URL specified in redirect_uri parameter.
// If authentication is successful, the following parameters are sent to redirect_uri:
//
// status: ok - successful authentication
// access_token - access token is passed in to all methods that require authentication
// expires_at - expiration date of access_token
// account_id - user ID
// nickname - user name.
//
// If authentication fails, the following parameters are sent to redirect_uri:
//
// status: error - authentication error
// code - error code
// message - error message.
//
// https://developers.wargaming.net/reference/all/wot/auth/login
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) AuthLogin(ctx context.Context, realm Realm, options *wot.AuthLoginOptions) (*wot.AuthLogin, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
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
		if options.Nofollow != nil {
			reqParam["nofollow"] = strconv.Itoa(*options.Nofollow)
		}
		if options.RedirectUri != nil {
			reqParam["redirect_uri"] = *options.RedirectUri
		}
	}

	var data *wot.AuthLogin
	err := service.client.getRequest(ctx, sectionWot, realm, "/auth/login/", reqParam, &data, nil)
	return data, err
}
