package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strconv"
	"strings"
)

// EncyclopediaPlanespecification returns information from Encyclopedia about aircraft specifications depending on installed modules.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planespecification
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// planeId:
//     Aircraft ID
func (service *WowpService) EncyclopediaPlanespecification(ctx context.Context, realm Realm, planeId int, options *wowp.EncyclopediaPlanespecificationOptions) (*wowp.EncyclopediaPlanespecification, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": strconv.Itoa(planeId),
	}
	if options != nil {
		if options.ModuleId != nil {
			reqParam["module_id"] = internal.SliceIntToString(options.ModuleId, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.BindId != nil {
			reqParam["bind_id"] = internal.SliceIntToString(options.BindId, ",")
		}
	}

	var data *wowp.EncyclopediaPlanespecification
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planespecification/", reqParam, &data)
	return data, err
}
