// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wows"
	"strings"
)

// ClansInfo returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wows/clans/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. Maximum limit: 100. Min value is 1.
func (service *WowsService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wows.ClansInfoOptions) (map[int]*wows.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
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

	var data map[int]*wows.ClansInfo
	err := service.client.getRequest(ctx, sectionWows, realm, "/clans/info/", reqParam, &data, nil)
	return data, err
}
