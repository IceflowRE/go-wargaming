// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotb"
	"strings"
)

// ClansGlossary returns information on clan entities.
//
// https://developers.wargaming.net/reference/all/wotb/clans/glossary
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WotbService) ClansGlossary(ctx context.Context, realm Realm, options *wotb.ClansGlossaryOptions) (*wotb.ClansGlossary, error) {
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

	var data *wotb.ClansGlossary
	err := service.client.getRequest(ctx, sectionWotb, realm, "/clans/glossary/", reqParam, &data, nil)
	return data, err
}
