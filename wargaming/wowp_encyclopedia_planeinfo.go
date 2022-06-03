package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/internal"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// EncyclopediaPlaneinfo returns aircraft details from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planeinfo
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
// planeId:
//     Aircraft ID. Maximum limit: 1000.
func (service *WowpService) EncyclopediaPlaneinfo(ctx context.Context, realm Realm, planeId []int, options *wowp.EncyclopediaPlaneinfoOptions) (*wowp.EncyclopediaPlaneinfo, error) {
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

	var data *wowp.EncyclopediaPlaneinfo
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planeinfo/", reqParam, &data)
	return data, err
}
