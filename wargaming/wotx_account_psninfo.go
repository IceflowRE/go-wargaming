package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotx"
	"strings"
)

// AccountPsninfo Method returns player data based on Play Station UID
//
// https://developers.wargaming.net/reference/all/wotx/account/psninfo
//
// realm:
//     Valid realms: RealmWgcb
// psnid:
//     Play Station UID. Maximum limit: 100.
func (service *wotxService) AccountPsninfo(ctx context.Context, realm Realm, psnid []string) ([]*wotx.AccountPsninfo, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"psnid": strings.Join(psnid, ","),
	}

	var data []*wotx.AccountPsninfo
	err := service.client.getRequest(ctx, sectionWotx, realm, "/account/psninfo/", reqParam, &data)
	return data, err
}
