// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
)

// ClanratingsDates returns dates with available rating data.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/dates
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) ClanratingsDates(ctx context.Context, realm Realm, options *wot.ClanratingsDatesOptions) (*wot.ClanratingsDates, *GenericMeta, error) {
	if !containsRealm([]Realm{RealmAsia, RealmEu, RealmNa}, realm) {
		return nil, nil, InvalidRealm{realm}
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
	}

	var data *wot.ClanratingsDates
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/dates/", reqParam, &data, &metaData)
	return data, metaData, err
}
