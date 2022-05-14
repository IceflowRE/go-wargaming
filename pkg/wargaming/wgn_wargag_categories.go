package wargaming

import (
	"strings"
	"strconv"
)

type WgnWargagCategories struct {
	// Content category ID
	CategoryId int `json:"category_id,omitempty"`
	// Content category name
	Name string `json:"name,omitempty"`
	// Content type
	Type_ string `json:"type,omitempty"`
}

// WgnWargagCategories Method returns information about content categories.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/categories
//
// category_id:
//     Content category ID
// type:
//     Content type. Valid values:
//     
//     "video" &mdash; Video content type 
//     "picture" &mdash; Image content type
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WgnWargagCategories(realm Realm, categoryId int, type_ string, fields []string) (*WgnWargagCategories, error) {
	if err := ValidateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"category_id": strconv.Itoa(categoryId),
		"type": type_,
		"fields": strings.Join(fields, ","),
	}

	var result *WgnWargagCategories
	err := client.doGetDataRequest(realm, "/wgn/wargag/categories/", reqParam, &result)
	return result, err
}
