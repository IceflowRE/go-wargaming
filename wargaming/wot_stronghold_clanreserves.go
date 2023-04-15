// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// StrongholdClanreserves returns information about available Reserves and their current status.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/clanreserves
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
func (service *WotService) StrongholdClanreserves(ctx context.Context, realm Realm, accessToken string, options *wot.StrongholdClanreservesOptions) ([]*wot.StrongholdClanreserves, error) {
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
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data []*wot.StrongholdClanreserves
	err := service.client.getRequest(ctx, sectionWot, realm, "/stronghold/clanreserves/", reqParam, &data, nil)
	return data, err
}
