package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// EncyclopediaPlanemodules returns information from Encyclopedia about modules available for specified aircrafts.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planemodules
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// planeId:
//     Aircraft ID. Maximum limit: 1000.
func (service *WowpService) EncyclopediaPlanemodules(ctx context.Context, realm Realm, planeId []int, options *wowp.EncyclopediaPlanemodulesOptions) (*wowp.EncyclopediaPlanemodules, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": internal.SliceIntToString(planeId, ","),
	}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = strings.Join(options.Type_, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wowp.EncyclopediaPlanemodules
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planemodules/", reqParam, &data)
	return data, err
}
