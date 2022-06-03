package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
	"strings"
)

// WargagTags returns information about content tags.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/tags
//
// realm:
//     Valid realms: RealmRu
func (service *WgnService) WargagTags(ctx context.Context, realm Realm, options *wgn.WargagTagsOptions) (*wgn.WargagTags, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{}
	if options != nil {
		if options.TagId != nil {
			reqParam["tag_id"] = strconv.Itoa(*options.TagId)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
	}

	var data *wgn.WargagTags
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wargag/tags/", reqParam, &data)
	return data, err
}
