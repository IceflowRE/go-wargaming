// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// AccountWtr returns player WTR rating
//
// https://developers.wargaming.net/reference/all/wot/account/wtr
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
//
// accountId:
//
//	Player ID. Maximum limit: 100.
func (service *WotService) AccountWtr(ctx context.Context, realm Realm, accountId []int, options *wot.AccountWtrOptions) (*wot.AccountWtr, error) {
	if !containsRealm([]Realm{RealmAsia, RealmEu, RealmNa}, realm) {
		return nil, InvalidRealm{realm}
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

	var data *wot.AccountWtr
	err := service.client.getRequest(ctx, sectionWot, realm, "/account/wtr/", reqParam, &data, nil)
	return data, err
}
