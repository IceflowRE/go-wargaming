// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEvents returns events information.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/events
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) GlobalmapEvents(ctx context.Context, realm Realm, options *wot.GlobalmapEventsOptions) ([]*wot.GlobalmapEvents, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.EventId != nil {
			reqParam["event_id"] = *options.EventId
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
		if options.Status != nil {
			reqParam["status"] = *options.Status
		}
	}

	var data []*wot.GlobalmapEvents
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/events/", reqParam, &data, nil)
	return data, err
}
