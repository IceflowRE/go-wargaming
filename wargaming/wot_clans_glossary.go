// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wot"
	"strings"
)

// ClansGlossary returns information on clan entities.
//
// https://developers.wargaming.net/reference/all/wot/clans/glossary
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotService) ClansGlossary(ctx context.Context, realm Realm, options *wot.ClansGlossaryOptions) (*wot.ClansGlossary, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
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

	var data *wot.ClansGlossary
	err := service.client.getRequest(ctx, sectionWot, realm, "/clans/glossary/", reqParam, &data, nil)
	return data, err
}
