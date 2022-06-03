package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strings"
)

// WgtvTags Method returns lists of game projects, categories, and programs.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/tags
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa, RealmRu
func (service *WgnService) WgtvTags(ctx context.Context, realm Realm, options *wgn.WgtvTagsOptions) (*wgn.WgtvTags, error) {
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

	var data *wgn.WgtvTags
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wgtv/tags/", reqParam, &data)
	return data, err
}
