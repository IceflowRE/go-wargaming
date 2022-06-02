package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEvents Method returns events information.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/events
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotService) GlobalmapEvents(ctx context.Context, realm Realm, options *wot.GlobalmapEventsOptions) ([]*wot.GlobalmapEvents, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Status != nil {
			reqParam["status"] = *options.Status
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.EventId != nil {
			reqParam["event_id"] = *options.EventId
		}
	}

	var data []*wot.GlobalmapEvents
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/events/", reqParam, &data)
	return data, err
}
