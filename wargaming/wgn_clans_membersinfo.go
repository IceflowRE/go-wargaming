// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// ClansMembersinfo returns detailed clan member information and brief clan details. Information is available for World of Tanks and World of Warplanes clans.This method will be removed. Use method Clan member details (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/membersinfo
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *WgnService) ClansMembersinfo(ctx context.Context, realm Realm, accountId []int, options *wgn.ClansMembersinfoOptions) (*wgn.ClansMembersinfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
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

	var data *wgn.ClansMembersinfo
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/membersinfo/", reqParam, &data, nil)
	return data, err
}
