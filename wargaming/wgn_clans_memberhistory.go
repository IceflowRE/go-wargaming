// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strconv"
	"strings"
)

// ClansMemberhistory returns information about player's clan history. Data on 10 last clan memberships are presented in the response.This method will be removed. Use method Player's clan history (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/memberhistory
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Account ID. Min value is 1.
func (service *WgnService) ClansMemberhistory(ctx context.Context, realm Realm, accountId int, options *wgn.ClansMemberhistoryOptions) (*wgn.ClansMemberhistory, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = *options.Game
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wgn.ClansMemberhistory
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/memberhistory/", reqParam, &data, nil)
	return data, err
}
