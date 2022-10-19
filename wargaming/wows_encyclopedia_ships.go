// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaShips returns list of ships available.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/ships
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaShips(ctx context.Context, realm Realm, options *wows.EncyclopediaShipsOptions) (*wows.EncyclopediaShips, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

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
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.ShipId != nil {
			reqParam["ship_id"] = internal.SliceIntToString(options.ShipId, ",")
		}
		if options.Type != nil {
			reqParam["type"] = strings.Join(options.Type, ",")
		}
	}

	var data *wows.EncyclopediaShips
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/ships/", reqParam, &data)
	return data, err
}
