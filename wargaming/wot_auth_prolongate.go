// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
)

// AuthProlongate generates new access_token based on the current token.
// This method is used when the player is still using the application but the current access_token is about to expire.
//
// https://developers.wargaming.net/reference/all/wot/auth/prolongate
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *WotService) AuthProlongate(ctx context.Context, realm Realm, accessToken string, options *wot.AuthProlongateOptions) (*wot.AuthProlongate, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}

	if options != nil {
		if options.ExpiresAt != nil {
			reqParam["expires_at"] = strconv.FormatInt(options.ExpiresAt.Unix(), 10)
		}
	}

	var data *wot.AuthProlongate
	err := service.client.postRequest(ctx, sectionWot, realm, "/auth/prolongate/", reqParam, &data, nil)
	return data, err
}
