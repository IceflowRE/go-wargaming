package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
	"strings"
)

// WgnWargagSearch Method searches content for specified text.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/search
//
// q:
//     Search text. Minimum length: 3 characters. Case insensitive.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// category_id:
//     Content category ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// rating_threshold:
//     Threshold of publication rating
// tag_id:
//     Tag ID
// type:
//     Content type. Valid values:
//     
//     "quote" &mdash; Quote content 
//     "video" &mdash; Video content 
//     "picture" &mdash; Image content type
func (client *Client) WgnWargagSearch(realm Realm, q string, accessToken string, categoryId int, fields []string, ratingThreshold int, tagId int, type_ string) (*wgn.WargagSearch, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"q": q,
		"access_token": accessToken,
		"category_id": strconv.Itoa(categoryId),
		"fields": strings.Join(fields, ","),
		"rating_threshold": strconv.Itoa(ratingThreshold),
		"tag_id": strconv.Itoa(tagId),
		"type": type_,
	}

	var result *wgn.WargagSearch
	err := client.doGetDataRequest(realm, "/wgn/wargag/search/", reqParam, &result)
	return result, err
}
