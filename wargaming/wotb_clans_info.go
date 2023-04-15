// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// ClansInfo returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wotb/clans/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. Maximum limit: 100. Min value is 1.
func (service *WotbService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wotb.ClansInfoOptions) (*wotb.ClansInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}

	if options != nil {
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wotb.ClansInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
