// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// EncyclopediaPlaneupgrades returns information from Encyclopedia about slots of aircrafts and lists of modules which are compatible with specified slots.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planeupgrades
//
// realm:
//
//	Valid realms: RealmEu, RealmNa
//
// planeId:
//
//	Aircraft ID. Maximum limit: 100.
func (service *WowpService) EncyclopediaPlaneupgrades(ctx context.Context, realm Realm, planeId []int, options *wowp.EncyclopediaPlaneupgradesOptions) (*wowp.EncyclopediaPlaneupgrades, *GenericMeta, error) {
	if !containsRealm([]Realm{RealmEu, RealmNa}, realm) {
		return nil, nil, InvalidRealm{realm}
	}

	reqParam := map[string]string{
		"plane_id": internal.SliceIntToString(planeId, ","),
	}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wowp.EncyclopediaPlaneupgrades
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planeupgrades/", reqParam, &data, &metaData)
	return data, metaData, err
}
