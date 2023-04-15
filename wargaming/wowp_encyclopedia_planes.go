// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wowp"
	"strings"
)

// EncyclopediaPlanes returns list of all aircrafts from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planes
//
// realm:
//     Valid realms: RealmEu, RealmNa
func (service *WowpService) EncyclopediaPlanes(ctx context.Context, realm Realm, options *wowp.EncyclopediaPlanesOptions) (*wowp.EncyclopediaPlanes, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa}); err != nil {
		return nil, nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.Type != nil {
			reqParam["type"] = strings.Join(options.Type, ",")
		}
	}

	var data *wowp.EncyclopediaPlanes
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planes/", reqParam, &data, &metaData)
	return data, metaData, err
}
