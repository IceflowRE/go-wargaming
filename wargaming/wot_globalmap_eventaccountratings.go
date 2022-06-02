package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventaccountratings Method returns account event rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountratings
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *wotService) GlobalmapEventaccountratings(ctx context.Context, realm Realm, eventId string, frontId string, options *wot.GlobalmapEventaccountratingsOptions) ([]*wot.GlobalmapEventaccountratings, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"event_id": eventId,
		"front_id": frontId,
	}
	if options != nil {
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.InRating != nil {
			reqParam["in_rating"] = strconv.Itoa(*options.InRating)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data []*wot.GlobalmapEventaccountratings
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventaccountratings/", reqParam, &data)
	return data, err
}
