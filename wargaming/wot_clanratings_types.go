// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wot"
)

// ClanratingsTypes returns details on ratings types and categories.
//
// https://developers.wargaming.net/reference/all/wot/clanratings/types
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) ClanratingsTypes(ctx context.Context, realm Realm) (*wot.ClanratingsTypes, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	var data *wot.ClanratingsTypes
	err := service.client.getRequest(ctx, sectionWot, realm, "/clanratings/types/", reqParam, &data, nil)
	return data, err
}
