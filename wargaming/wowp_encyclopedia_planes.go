package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wowp"
	"strings"
)

// EncyclopediaPlanes Method returns list of all aircrafts from Encyclopedia.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planes
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
func (service *WowpService) EncyclopediaPlanes(ctx context.Context, realm Realm, options *wowp.EncyclopediaPlanesOptions) (map[string]*wowp.EncyclopediaPlanes, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = strings.Join(options.Type_, ",")
		}
		if options.Nation != nil {
			reqParam["nation"] = strings.Join(options.Nation, ",")
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data map[string]*wowp.EncyclopediaPlanes
	err := service.client.getRequest(ctx, sectionWowp, realm, "/encyclopedia/planes/", reqParam, &data)
	return data, err
}
