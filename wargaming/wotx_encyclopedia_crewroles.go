package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotx"
	"strings"
)

// EncyclopediaCrewroles Method returns information about crews.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/crewroles
//
// realm:
//     Valid realms: RealmWgcb
func (service *wotxService) EncyclopediaCrewroles(ctx context.Context, realm Realm, options *wotx.EncyclopediaCrewrolesOptions) (*wotx.EncyclopediaCrewroles, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotx.EncyclopediaCrewroles
	err := service.client.getRequest(ctx, sectionWotx, realm, "/encyclopedia/crewroles/", reqParam, &data)
	return data, err
}
