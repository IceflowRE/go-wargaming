package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// EncyclopediaPlaneupgrades Method returns information from Encyclopedia about slots of aircrafts and lists of modules which are compatible with specified slots.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planeupgrades
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// planeId:
//     Aircraft ID. Maximum limit: 100.
func (service *wowpService) EncyclopediaPlaneupgrades(ctx context.Context, realm Realm, planeId []int, options *wowp.EncyclopediaPlaneupgradesOptions) (*wowp.EncyclopediaPlaneupgrades, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": internal.SliceIntToString(planeId, ","),
	}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wowp.EncyclopediaPlaneupgrades
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planeupgrades/", reqParam, &data)
	return data, err
}
