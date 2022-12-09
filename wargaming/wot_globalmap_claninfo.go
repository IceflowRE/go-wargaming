// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strings"
)

// GlobalmapClaninfo returns clan data on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/claninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
func (service *WotService) GlobalmapClaninfo(ctx context.Context, realm Realm, clanId []int, options *wot.GlobalmapClaninfoOptions) (*wot.GlobalmapClaninfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.GlobalmapClaninfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/claninfo/", reqParam, &data, nil)
	return data, err
}
