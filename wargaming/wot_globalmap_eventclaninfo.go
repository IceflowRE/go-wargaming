// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventclaninfo returns clan's statistics for a specific event.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventclaninfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method. Maximum limit: 10.
func (service *WotService) GlobalmapEventclaninfo(ctx context.Context, realm Realm, clanId int, eventId string, frontId []string, options *wot.GlobalmapEventclaninfoOptions) (*wot.GlobalmapEventclaninfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"clan_id":  strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": strings.Join(frontId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.GlobalmapEventclaninfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventclaninfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
