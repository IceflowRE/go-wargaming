// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// ClansInfo returns detailed clan information. Information is available for World of Tanks and World of Warplanes clans.This method will be removed. Use method Clan details (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/info
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. Maximum limit: 100.
func (service *WgnService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wgn.ClansInfoOptions) (*wgn.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}

	if options != nil {
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = *options.Game
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.MembersKey != nil {
			reqParam["members_key"] = *options.MembersKey
		}
	}

	var data *wgn.ClansInfo
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/info/", reqParam, &data, nil)
	return data, err
}
