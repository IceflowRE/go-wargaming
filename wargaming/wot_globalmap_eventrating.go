// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventrating returns event clan rating
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventrating
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *WotService) GlobalmapEventrating(ctx context.Context, realm Realm, eventId string, frontId string, options *wot.GlobalmapEventratingOptions) ([]*wot.GlobalmapEventrating, error) {
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
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.GlobalmapEventrating
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventrating/", reqParam, &data, nil)
	return data, err
}
