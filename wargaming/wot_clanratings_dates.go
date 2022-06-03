package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
)

// ClanratingsDates Method returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/dates
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WotService) ClanratingsDates(ctx context.Context, realm Realm, options *wot.ClanratingsDatesOptions) (*wot.ClanratingsDates, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
	}

	var data *wot.ClanratingsDates
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/dates/", reqParam, &data)
	return data, err
}
