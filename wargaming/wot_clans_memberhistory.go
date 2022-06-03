package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strconv"
	"strings"
)

// ClansMemberhistory Method returns information about player's clan history. Data on 10 last clan memberships are presented in the response.This method will be removed. Use method Player's clan history (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wot/clans/memberhistory
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Account ID. Min value is 1.
func (service *WotService) ClansMemberhistory(ctx context.Context, realm Realm, accountId int, options *wot.ClansMemberhistoryOptions) (*wot.ClansMemberhistory, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.ClansMemberhistory
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/memberhistory/", reqParam, &data)
	return data, err
}
