// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// ClanratingsTop returns the list of top clans by specified parameters.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/top
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// rankField:
//     Rating category
func (service *WotService) ClanratingsTop(ctx context.Context, realm Realm, rankField string, options *wot.ClanratingsTopOptions) ([]*wot.ClanratingsTop, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
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
		if options.PageNo != nil {
			reqParam["page_no"] = strconv.Itoa(*options.PageNo)
		}
	}

	var data []*wot.ClanratingsTop
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/top/", reqParam, &data, &metaData)
	return data, metaData, err
}
