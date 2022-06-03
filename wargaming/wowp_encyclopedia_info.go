package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
)

// EncyclopediaInfo Method returns information about Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/info
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
func (service *WowpService) EncyclopediaInfo(ctx context.Context, realm Realm) (*wowp.EncyclopediaInfo, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	var data *wowp.EncyclopediaInfo
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/info/", reqParam, &data)
	return data, err
}
