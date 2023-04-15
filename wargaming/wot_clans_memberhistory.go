// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strconv"
	"strings"
)

// ClansMemberhistory returns information about player's clan history. Data on 10 last clan memberships are presented in the response.This method will be removed. Use method Player's clan history (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wot/clans/memberhistory
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Account ID. Min value is 1.
func (service *WotService) ClansMemberhistory(ctx context.Context, realm Realm, accountId int, options *wot.ClansMemberhistoryOptions) (*wot.ClansMemberhistory, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wot.ClansMemberhistory
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/memberhistory/", reqParam, &data, &metaData)
	return data, metaData, err
}
