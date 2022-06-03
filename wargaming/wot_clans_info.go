package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// ClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wot/clans/info
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// clanId:
//     Clan ID. Maximum limit: 100.
func (service *wotService) ClansInfo(ctx context.Context, realm Realm, clanId []int, options *wot.ClansInfoOptions) (*wot.ClansInfo, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": internal.SliceIntToString(clanId, ","),
	}
	if options != nil {
		if options.MembersKey != nil {
			reqParam["members_key"] = *options.MembersKey
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Extra != nil {
			reqParam["extra"] = strings.Join(options.Extra, ",")
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wot.ClansInfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/info/", reqParam, &data)
	return data, err
}
