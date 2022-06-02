package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strings"
)

// GlobalmapClanprovinces Method returns lists of clans provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/clanprovinces
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     List of clan IDs. To get a clan ID, use the Clans method. Maximum limit: 10.
func (service *wotService) GlobalmapClanprovinces(ctx context.Context, realm Realm, clanId []int, options *wot.GlobalmapClanprovincesOptions) (*wot.GlobalmapClanprovinces, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wot.GlobalmapClanprovinces
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/clanprovinces/", reqParam, &data)
	return data, err
}
