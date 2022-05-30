package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
	"strings"
)

// WgnWargagCategories Method returns information about content categories.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/categories
//
// categoryId:
//     Content category ID
// type_:
//     Content type. Valid values:
//     
//     "video" - Video content type 
//     "picture" - Image content type
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WgnWargagCategories(realm Realm, categoryId int, type_ string, fields []string) (*wgn.WargagCategories, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"category_id": strconv.Itoa(categoryId),
		"type": type_,
		"fields": strings.Join(fields, ","),
	}

	var result *wgn.WargagCategories
	err := client.doGetDataRequest(realm, "/wgn/wargag/categories/", reqParam, &result)
	return result, err
}
