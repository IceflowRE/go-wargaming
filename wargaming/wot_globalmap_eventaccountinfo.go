// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventaccountinfo returns player's statistics for a specific event
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountinfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method. Maximum limit: 10.
func (service *WotService) GlobalmapEventaccountinfo(ctx context.Context, realm Realm, eventId string, frontId []string, options *wot.GlobalmapEventaccountinfoOptions) (*wot.GlobalmapEventaccountinfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"event_id": eventId,
		"front_id": strings.Join(frontId, ","),
	}

	if options != nil {
		if options.ClanId != nil {
			reqParam["clan_id"] = strconv.Itoa(*options.ClanId)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.AccountId != nil {
			reqParam["account_id"] = strconv.Itoa(*options.AccountId)
		}
	}

	var data *wot.GlobalmapEventaccountinfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventaccountinfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
