// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// EncyclopediaPlaneinfo returns aircraft details from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planeinfo
//
// realm:
//     Valid realms: RealmEu, RealmNa
// planeId:
//     Aircraft ID. Maximum limit: 1000.
func (service *WowpService) EncyclopediaPlaneinfo(ctx context.Context, realm Realm, planeId []int, options *wowp.EncyclopediaPlaneinfoOptions) (*wowp.EncyclopediaPlaneinfo, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
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

	var data *wowp.EncyclopediaPlaneinfo
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planeinfo/", reqParam, &data, &metaData)
	return data, metaData, err
}
