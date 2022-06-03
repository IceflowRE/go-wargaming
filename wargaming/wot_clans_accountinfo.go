package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wot"
	"strings"
)

// ClansAccountinfo Method returns detailed clan member information and brief clan details.
//
// https://developers.wargaming.net/reference/all/wot/clans/accountinfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *WotService) ClansAccountinfo(ctx context.Context, realm Realm, accountId []int, options *wot.ClansAccountinfoOptions) (*wot.ClansAccountinfo, error) {
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
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wot.ClansAccountinfo
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/accountinfo/", reqParam, &data)
	return data, err
}
