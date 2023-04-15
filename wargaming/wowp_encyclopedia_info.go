// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
)

// EncyclopediaInfo returns information about Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/info
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) EncyclopediaInfo(ctx context.Context, realm Realm) (*wowp.EncyclopediaInfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	var data *wowp.EncyclopediaInfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/info/", reqParam, &data, &metaData)
	return data, metaData, err
}
