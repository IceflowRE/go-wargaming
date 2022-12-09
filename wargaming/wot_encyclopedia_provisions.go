// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strconv"
	"strings"
)

// EncyclopediaProvisions returns a list of available equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/provisions
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaProvisions(ctx context.Context, realm Realm, options *wot.EncyclopediaProvisionsOptions) (*wot.EncyclopediaProvisions, error) {
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
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.ProvisionId != nil {
			reqParam["provision_id"] = internal.SliceIntToString(options.ProvisionId, ",")
		}
		if options.Type != nil {
			reqParam["type"] = strings.Join(options.Type, ",")
		}
	}

	var data *wot.EncyclopediaProvisions
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/provisions/", reqParam, &data, nil)
	return data, err
}
