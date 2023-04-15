// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// EncyclopediaPersonalmissions returns details on Personal Missions on the basis of specified campaign IDs, operation IDs, mission branch and tag IDs.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/personalmissions
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) EncyclopediaPersonalmissions(ctx context.Context, realm Realm, options *wot.EncyclopediaPersonalmissionsOptions) (*wot.EncyclopediaPersonalmissions, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.CampaignId != nil {
			reqParam["campaign_id"] = internal.SliceIntToString(options.CampaignId, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.OperationId != nil {
			reqParam["operation_id"] = internal.SliceIntToString(options.OperationId, ",")
		}
		if options.SetId != nil {
			reqParam["set_id"] = internal.SliceIntToString(options.SetId, ",")
		}
		if options.Tag != nil {
			reqParam["tag"] = strings.Join(options.Tag, ",")
		}
	}

	var data *wot.EncyclopediaPersonalmissions
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/encyclopedia/personalmissions/", reqParam, &data, &metaData)
	return data, metaData, err
}
