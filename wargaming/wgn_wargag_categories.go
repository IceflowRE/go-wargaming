package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
	"strings"
)

// WargagCategories returns information about content categories.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/categories
//
// realm:
//     Valid realms: RealmRu
// type_:
//     Content type. Valid values:
//
//     "video" - Video content type
//     "picture" - Image content type
func (service *WgnService) WargagCategories(ctx context.Context, realm Realm, type_ string, options *wgn.WargagCategoriesOptions) (*wgn.WargagCategories, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"type": type_,
	}
	if options != nil {
		if options.Fields != nil {
			reqParam["fields"] = strings.Join(options.Fields, ",")
		}
		if options.CategoryId != nil {
			reqParam["category_id"] = strconv.Itoa(*options.CategoryId)
		}
	}

	var data *wgn.WargagCategories
	err := service.client.getRequest(ctx, sectionWgn, realm, "/wargag/categories/", reqParam, &data)
	return data, err
}
