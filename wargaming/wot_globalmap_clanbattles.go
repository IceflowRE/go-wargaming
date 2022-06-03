package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapClanbattles returns list of clan's battles on the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/clanbattles
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID. To get a clan ID, use the Clans method.
func (service *WotService) GlobalmapClanbattles(ctx context.Context, realm Realm, clanId int, options *wot.GlobalmapClanbattlesOptions) ([]*wot.GlobalmapClanbattles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": strconv.Itoa(clanId),
	}
	if options != nil {
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wot.GlobalmapClanbattles
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/clanbattles/", reqParam, &data)
	return data, err
}
