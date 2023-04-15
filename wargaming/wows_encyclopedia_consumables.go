// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wows"
	"strconv"
	"strings"
)

// EncyclopediaConsumables returns information about consumables: camouflages, flags, and upgrades.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/consumables
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WowsService) EncyclopediaConsumables(ctx context.Context, realm Realm, options *wows.EncyclopediaConsumablesOptions) (*wows.EncyclopediaConsumables, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.ConsumableId != nil {
			reqParam["consumable_id"] = internal.SliceIntToString(options.ConsumableId, ",")
		}
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
		if options.Type != nil {
			reqParam["type"] = *options.Type
		}
	}

	var data *wows.EncyclopediaConsumables
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWows, realm, "/encyclopedia/consumables/", reqParam, &data, &metaData)
	return data, metaData, err
}
