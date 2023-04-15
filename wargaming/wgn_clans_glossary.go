// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// ClansGlossary returns information on clan entities in World of Tanks and World of Warplanes.This method will be removed. Use method Clan glossary (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/glossary
//
// Deprecated: Attention! The method is deprecated.
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WgnService) ClansGlossary(ctx context.Context, realm Realm, options *wgn.ClansGlossaryOptions) (*wgn.ClansGlossary, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}

	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.Game != nil {
			reqParam["game"] = *options.Game
		}
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
	}

	var data *wgn.ClansGlossary
	err := service.client.getRequest(ctx, sectionWgn, realm, "/clans/glossary/", reqParam, &data, nil)
	return data, err
}
