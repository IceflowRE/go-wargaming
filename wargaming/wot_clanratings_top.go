package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// ClanratingsTop Method returns the list of top clans by specified parameters.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/top
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// rankField:
//     Rating category
func (service *wotService) ClanratingsTop(ctx context.Context, realm Realm, rankField string, options *wot.ClanratingsTopOptions) ([]*wot.ClanratingsTop, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"rank_field": rankField,
	}
	if options != nil {
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
		if options.Date != nil {
			reqParam["date"] = strconv.FormatInt(options.Date.Unix(), 10)
		}
	}

	var data []*wot.ClanratingsTop
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/top/", reqParam, &data)
	return data, err
}
