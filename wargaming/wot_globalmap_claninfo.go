package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strings"
)

// GlobalmapClaninfo Method returns clan data on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/claninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Maximum limit: 10.
func (service *wotService) GlobalmapClaninfo(ctx context.Context, realm Realm, clanId []int, options *wot.GlobalmapClaninfoOptions) (*wot.GlobalmapClaninfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}
	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wot.GlobalmapClaninfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/claninfo/", reqParam, &data)
	return data, err
}
