package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// ClanratingsNeighbors Method returns list of adjacent positions in specified clan rating.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/neighbors
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID
// rankField:
//     Rating category
func (service *wotService) ClanratingsNeighbors(ctx context.Context, realm Realm, clanId int, rankField string, options *wot.ClanratingsNeighborsOptions) ([]*wot.ClanratingsNeighbors, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id":    strconv.Itoa(clanId),
		"rank_field": rankField,
	}
	if options != nil {
		if options.Limit != nil {
			reqParam["limit"] = strconv.Itoa(*options.Limit)
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
	}

	var data []*wot.ClanratingsNeighbors
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/neighbors/", reqParam, &data)
	return data, err
}
