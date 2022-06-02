package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wowp"
	"strings"
)

// ClansGlossary Method returns information on clan entities.
//
// https://developers.wargaming.net/reference/all/wowp/clans/glossary
//
// realm:
//     Valid realms: RealmEu, RealmNa, RealmRu
func (service *wowpService) ClansGlossary(ctx context.Context, realm Realm, options *wowp.ClansGlossaryOptions) (*wowp.ClansGlossary, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
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

	var data *wowp.ClansGlossary
	err := service.client.getRequest(ctx, sectionWowp, realm, "/clans/glossary/", reqParam, &data)
	return data, err
}
