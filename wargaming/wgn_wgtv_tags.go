// Auto generated file!

package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgn"
	"strings"
)

// WgtvTags returns lists of game projects, categories, and programs.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/tags
//
// realm:
//     Valid realms: RealmAsia, RealmEu, RealmNa
func (service *WgnService) WgtvTags(ctx context.Context, realm Realm, options *wgn.WgtvTagsOptions) (*wgn.WgtvTags, *GenericMeta, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return nil, nil, err
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

	var data *wgn.WgtvTags
	var metaData *GenericMeta
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wgtv/tags/", reqParam, &data, &metaData)
	return data, metaData, err
}
