// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wotx"
	"strings"
)

// ClansGlossary returns general information about clans.
//
// https://developers.wargaming.net/reference/all/wotx/clans/glossary
//
// realm:
//     Valid realms: RealmWgcb
func (service *WotxService) ClansGlossary(ctx context.Context, realm Realm, options *wotx.ClansGlossaryOptions) (*wotx.ClansGlossary, error) {
	if err := validateRealm(realm, []Realm{RealmWgcb}); err != nil {
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

	var data *wotx.ClansGlossary
	err := service.client.getRequest(ctx, sectionWotx, realm, "/clans/glossary/", reqParam, &data, nil)
	return data, err
}
