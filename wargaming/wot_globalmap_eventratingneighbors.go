package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventratingneighbors Method returns list of adjacent positions in event clan rating
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventratingneighbors
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID. To get a clan ID, use the Clans method. Min value is 1.
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *wotService) GlobalmapEventratingneighbors(ctx context.Context, realm Realm, clanId int, eventId string, frontId string, options *wot.GlobalmapEventratingneighborsOptions) ([]*wot.GlobalmapEventratingneighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id":  strconv.Itoa(clanId),
		"event_id": eventId,
		"front_id": frontId,
	}
	if options != nil {
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wot.GlobalmapEventratingneighbors
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventratingneighbors/", reqParam, &data)
	return data, err
}