package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strings"
)

// ClansMembersinfo Method returns detailed clan member information and brief clan details. Information is available for World of Tanks and World of Warplanes clans.This method will be removed. Use method Clan member details (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/membersinfo
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *wgnService) ClansMembersinfo(ctx context.Context, realm Realm, accountId []int, options *wgn.ClansMembersinfoOptions) (*wgn.ClansMembersinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Game != nil {
			reqParam["game"] = *options.Game
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wgn.ClansMembersinfo
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/membersinfo/", reqParam, &data)
	return data, err
}
