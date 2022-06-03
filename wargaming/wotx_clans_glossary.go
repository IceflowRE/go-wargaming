package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wotx"
	"strings"
)

// ClansGlossary Method returns general information about clans.
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
		if options.Language != nil {
			reqParam["language"] = *options.Language
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wotx.ClansGlossary
	err := service.client.getRequest(ctx, sectionWotx, realm, "/clans/glossary/", reqParam, &data)
	return data, err
}
