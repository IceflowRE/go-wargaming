// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wowp"
	"strings"
)

// ClansInfo returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wowp/clans/info
//
// realm:
//     Valid realms: RealmEu, RealmNa
// clanId:
//     Clan ID. Maximum limit: 100. Min value is 1.
func (service *WowpService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wowp.ClansInfoOptions) (*wowp.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, err
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

	var data *wowp.ClansInfo
	err := service.client.getRequest(ctx, sectionWowp, realm, "/clans/info/", reqParam, &data, nil)
	return data, err
}
