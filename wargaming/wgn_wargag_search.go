package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
	"strings"
)

// WargagSearch Method searches content for specified text.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/search
//
// realm:
//     Valid realms: RealmRu
// q:
//     Search text. Minimum length: 3 characters. Case insensitive.
func (service *WgnService) WargagSearch(ctx context.Context, realm Realm, q string, options *wgn.WargagSearchOptions) (*wgn.WargagSearch, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"q": q,
	}
	if options != nil {
		if options.Type_ != nil {
			reqParam["type"] = *options.Type_
		}
		if options.TagId != nil {
			reqParam["tag_id"] = strconv.Itoa(*options.TagId)
		}
		if options.RatingThreshold != nil {
			reqParam["rating_threshold"] = strconv.Itoa(*options.RatingThreshold)
		}
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.CategoryId != nil {
			reqParam["category_id"] = strconv.Itoa(*options.CategoryId)
		}
		if options.AccessToken != nil {
			reqParam["access_token"] = *options.AccessToken
		}
	}

	var data *wgn.WargagSearch
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wargag/search/", reqParam, &data)
	return data, err
}
