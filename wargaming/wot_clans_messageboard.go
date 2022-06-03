package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// ClansMessageboard Method returns messages of clan message board.This method will be removed. Use method Message board (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wot/clans/messageboard
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *wotService) ClansMessageboard(ctx context.Context, realm Realm, accessToken string, options *wot.ClansMessageboardOptions) (map[string][]*wot.ClansMessageboard, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
	}
	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data map[string][]*wot.ClansMessageboard
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/messageboard/", reqParam, &data)
	return data, err
}
