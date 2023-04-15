// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strconv"
	"strings"
)

// ClansList searches and sorts clans by the following logic:
//
// exact match of clan tag placed first
// exact match of clan name placed second
// name or tag matches are placed next
// all comparisons are performed case insensitive
// expression [wg] is converted to wg and considered as the clan tag
// search for expression "[wg] clan" is performed by exact match of clan name and tag
//
// Disbanded, NPC and technically frozen clans are excluded from response. Search is executed across World of Tanks and World of Warplanes.This method will be removed. Use method Clans (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/list
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WgnService) ClansList(ctx context.Context, realm Realm, options *wgn.ClansListOptions) (*wgn.ClansList, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = strings.Join(options.Game, ",")
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
		if options.Search != nil {
			reqParam["search"] = *options.Search
		}
	}

	var data *wgn.ClansList
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/list/", reqParam, &data, nil)
	return data, err
}
