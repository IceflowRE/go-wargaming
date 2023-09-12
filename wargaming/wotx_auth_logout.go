// Auto generated file!

package wargaming

import (
	"context"
)

// AuthLogout deletes user's access_token.
// After this method is called, access_token becomes invalid.
//
// https://developers.wargaming.net/reference/all/wotx/auth/logout
//
// realm:
//
//	Valid realms: RealmWgcb
//
// accessToken:
//
//	Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *WotxService) AuthLogout(ctx context.Context, realm Realm, accessToken string) error {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}

	err := service.client.postRequest(ctx, sectionWotx, realm, "/auth/logout/", reqParam, nil, nil)
	return err
}
