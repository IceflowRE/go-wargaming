package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaConsumables Method returns information about consumables: camouflages, flags, and upgrades.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/consumables
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wowsService) EncyclopediaConsumables(ctx context.Context, realm Realm, options *wows.EncyclopediaConsumablesOptions) (*wows.EncyclopediaConsumables, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
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
		if options.ConsumableId != nil {
			reqParam["consumable_id"] = internal.SliceIntToString(options.ConsumableId, ",")
		}
	}

	var data *wows.EncyclopediaConsumables
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/consumables/", reqParam, &data)
	return data, err
}
