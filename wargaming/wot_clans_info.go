// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// ClansInfo returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wot/clans/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// clanId:
//     Clan ID. Maximum limit: 100.
func (service *WotService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wot.ClansInfoOptions) (*wot.ClansInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
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
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.MembersKey != nil {
			reqParam["members_key"] = *options.MembersKey
		}
	}

	var data *wot.ClansInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
