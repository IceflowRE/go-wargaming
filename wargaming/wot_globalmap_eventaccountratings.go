// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventaccountratings returns account event rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountratings
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *WotService) GlobalmapEventaccountratings(ctx context.Context, realm Realm, eventId string, frontId string, options *wot.GlobalmapEventaccountratingsOptions) ([]*wot.GlobalmapEventaccountratings, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"event_id": eventId,
		"front_id": frontId,
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.InRating != nil {
			reqParam["in_rating"] = strconv.Itoa(*options.InRating)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.GlobalmapEventaccountratings
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventaccountratings/", reqParam, &data, nil)
	return data, err
}
