package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotb"
	"strings"
)

// ClansInfo returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wotb/clans/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID. Maximum limit: 100. Min value is 1.
func (service *WotbService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wotb.ClansInfoOptions) (*wotb.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
	}

	var data *wotb.ClansInfo
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/info/", reqParam, &data)
	return data, err
}
