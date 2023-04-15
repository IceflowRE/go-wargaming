// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// ClansAccountinfo returns detailed clan member information and brief clan details.
//
// https://developers.wargaming.net/reference/all/wot/clans/accountinfo
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
// accountId:
//     Account ID. Maximum limit: 100. Min value is 1.
func (service *WotService) ClansAccountinfo(ctx context.Context, realm Realm, accountId []int, options *wot.ClansAccountinfoOptions) (*wot.ClansAccountinfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{
		"account_id": internal.SliceIntToString(accountId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wot.ClansAccountinfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/accountinfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
