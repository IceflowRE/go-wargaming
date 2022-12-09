// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strings"
)

// GlobalmapClanprovinces returns lists of clans provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/clanprovinces
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     List of clan IDs. To get a clan ID, use the Clans method. Maximum limit: 10.
func (service *WotService) GlobalmapClanprovinces(ctx context.Context, realm Realm, clanId []int, options *wot.GlobalmapClanprovincesOptions) (*wot.GlobalmapClanprovinces, error) {
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
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wot.GlobalmapClanprovinces
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/clanprovinces/", reqParam, &data, nil)
	return data, err
}
