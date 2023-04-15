// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// ClansMessageboard returns messages of clan message board.This method will be removed. Use method Message board (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/messageboard
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *WgnService) ClansMessageboard(ctx context.Context, realm Realm, accessToken string, options *wgn.ClansMessageboardOptions) (*wgn.ClansMessageboard, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = *options.Game
		}
	}

	var data *wgn.ClansMessageboard
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/messageboard/", reqParam, &data, nil)
	return data, err
}
