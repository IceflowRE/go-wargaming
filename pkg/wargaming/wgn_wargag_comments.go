package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
	"strings"
)

// WgnWargagComments Method returns comments to the content.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/comments
//
// contentId:
//     Publication ID
//     Parameter is required.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// pageNo:
//     Result page number
func (client *Client) WgnWargagComments(realm Realm, contentId int, fields []string, pageNo int) (*wgn.WargagComments, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"content_id": strconv.Itoa(contentId),
		"fields": strings.Join(fields, ","),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *wgn.WargagComments
	err := client.doGetDataRequest(realm, "/wgn/wargag/comments/", reqParam, &result)
	return result, err
}
