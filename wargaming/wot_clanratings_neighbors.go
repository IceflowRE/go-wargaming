// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// ClanratingsNeighbors returns list of adjacent positions in specified clan rating.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/neighbors
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID
// rankField:
//     Rating category
func (service *WotService) ClanratingsNeighbors(ctx context.Context, realm Realm, clanId int, rankField string, options *wot.ClanratingsNeighborsOptions) ([]*wot.ClanratingsNeighbors, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"clan_id":    strconv.Itoa(clanId),
		"rank_field": rankField,
	}

	if options != nil {
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
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
	}

	var data []*wot.ClanratingsNeighbors
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/neighbors/", reqParam, &data, &metaData)
	return data, metaData, err
}
