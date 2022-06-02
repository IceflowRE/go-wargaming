package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/wargaming/wotb"
	"strings"
)

// ClansGlossary Method returns information on clan entities.
//
// https://developers.wargaming.net/reference/all/wotb/clans/glossary
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *wotbService) ClansGlossary(ctx context.Context, realm Realm, options *wotb.ClansGlossaryOptions) (*wotb.ClansGlossary, error) {
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

	var data *wotb.ClansGlossary
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/glossary/", reqParam, &data)
	return data, err
}
