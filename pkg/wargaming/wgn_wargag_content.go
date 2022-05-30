package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"strconv"
	"strings"
)

// WgnWargagContent Method returns information about content.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/content
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// accountId:
//     Publication author ID
// categoryId:
//     Content category ID
// contentId:
//     Content ID. When this parameter is specified, all other passed-in parameters are ignored.
// date:
//     Publication date
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// orderBy:
//     Sorting. Default is "-date". Valid values:
//     
//     "date" - by publication date 
//     "-date" - by publication date in reverse order (by default)
//     "rating" - by rating value 
//     "-rating" - by rating value in reverse order
// pageNo:
//     Result page number
// ratingThreshold:
//     Threshold of publication rating
// tagId:
//     Tag ID
// type_:
//     Content type. Valid values:
//     
//     "quote" - Quote content 
//     "video" - Video content 
//     "picture" - Image content type
func (client *Client) WgnWargagContent(realm Realm, accessToken string, accountId int, categoryId int, contentId int, date wgnTime.UnixTime, fields []string, orderBy string, pageNo int, ratingThreshold int, tagId int, type_ string) (*wgn.WargagContent, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"account_id": strconv.Itoa(accountId),
		"category_id": strconv.Itoa(categoryId),
		"content_id": strconv.Itoa(contentId),
		"date": strconv.FormatInt(date.Unix(), 10),
		"fields": strings.Join(fields, ","),
		"order_by": orderBy,
		"page_no": strconv.Itoa(pageNo),
		"rating_threshold": strconv.Itoa(ratingThreshold),
		"tag_id": strconv.Itoa(tagId),
		"type": type_,
	}

	var result *wgn.WargagContent
	err := client.doGetDataRequest(realm, "/wgn/wargag/content/", reqParam, &result)
	return result, err
}