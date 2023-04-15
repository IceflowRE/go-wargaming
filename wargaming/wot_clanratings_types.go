// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
)

// ClanratingsTypes returns details on ratings types and categories.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/types
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) ClanratingsTypes(ctx context.Context, realm Realm) (*wot.ClanratingsTypes, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	var data *wot.ClanratingsTypes
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/types/", reqParam, &data, &metaData)
	return data, metaData, err
}
