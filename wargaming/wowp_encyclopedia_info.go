// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wowp"
)

// EncyclopediaInfo returns information about Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/info
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) EncyclopediaInfo(ctx context.Context, realm Realm) (*wowp.EncyclopediaInfo, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	var data *wowp.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/info/", reqParam, &data, nil)
	return data, err
}
