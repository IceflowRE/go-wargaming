// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapClanbattles returns list of clan's battles on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/clanbattles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. To get a clan ID, use the Clans method.
func (service *WotService) GlobalmapClanbattles(ctx context.Context, realm Realm, clanId int, options *wot.GlobalmapClanbattlesOptions) ([]*wot.GlobalmapClanbattles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.GlobalmapClanbattles
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/clanbattles/", reqParam, &data, nil)
	return data, err
}
