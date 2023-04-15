// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strings"
)

// EncyclopediaCrewroles returns information about crews.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/crewroles
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) EncyclopediaCrewroles(ctx context.Context, realm Realm, options *wotx.EncyclopediaCrewrolesOptions) (*wotx.EncyclopediaCrewroles, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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
	}

	var data *wotx.EncyclopediaCrewroles
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/crewroles/", reqParam, &data, &metaData)
	return data, metaData, err
}
