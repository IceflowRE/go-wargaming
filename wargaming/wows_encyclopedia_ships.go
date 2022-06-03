package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaShips Method returns list of ships available.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/ships
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) EncyclopediaShips(ctx context.Context, realm Realm, options *wows.EncyclopediaShipsOptions) (*wows.EncyclopediaShips, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = strings.Join(options.Type_, ",")
		}
		if options.ShipId != nil {
			reqParam["ship_id"] = internal.SliceIntToString(options.ShipId, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
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

	var data *wows.EncyclopediaShips
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/ships/", reqParam, &data)
	return data, err
}
