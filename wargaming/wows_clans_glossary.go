package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wows"
	"strings"
)

// ClansGlossary Method returns information on clan entities.
//
// https://developers.wargaming.net/reference/all/wows/clans/glossary
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WowsService) ClansGlossary(ctx context.Context, realm Realm, options *wows.ClansGlossaryOptions) (*wows.ClansGlossary, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
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

	var data *wows.ClansGlossary
	err := service.client.getRequest(ctx, sectionWows, realm, "/clans/glossary/", reqParam, &data)
	return data, err
}
