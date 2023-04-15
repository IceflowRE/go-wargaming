// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// GlobalmapEventaccountratingneighbors returns adjacent position in account event rating.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/eventaccountratingneighbors
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Account ID. Min value is 1.
// eventId:
//     Event ID. To get an event ID, use the Events method.
// frontId:
//     Front ID. To get a front ID, use the Fronts method.
func (service *WotService) GlobalmapEventaccountratingneighbors(ctx context.Context, realm Realm, accountId int, eventId string, frontId string, options *wot.GlobalmapEventaccountratingneighborsOptions) ([]*wot.GlobalmapEventaccountratingneighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"event_id":   eventId,
		"front_id":   frontId,
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.NeighboursCount != nil {
			reqParam["neighbours_count"] = strconv.Itoa(*options.NeighboursCount)
		}
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.GlobalmapEventaccountratingneighbors
	err := service.client.getRequest(ctx, sectionWot, realm, "/globalmap/eventaccountratingneighbors/", reqParam, &data, nil)
	return data, err
}
