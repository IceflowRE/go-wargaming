package wargaming

import (
	"strconv"
	"strings"
)

type WgnWargagTags struct {
	// Tag name
	Name string `json:"name,omitempty"`
	// Tag ID
	TagId int `json:"tag_id,omitempty"`
}

// WgnWargagTags Method returns information about content tags.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/tags
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// tag_id:
//     Tag ID
func (client *Client) WgnWargagTags(realm Realm, fields []string, tagId int) (*WgnWargagTags, error) {
	if err := ValidateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"tag_id": strconv.Itoa(tagId),
	}

	var result *WgnWargagTags
	err := client.doGetDataRequest(realm, "/wgn/wargag/tags/", reqParam, &result)
	return result, err
}
