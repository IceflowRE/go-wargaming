package wargaming

import (
	"strings"
	"strconv"
)

type WgnWargagComments struct {
	// Comment author ID
	AccountId int `json:"account_id,omitempty"`
	// Comment text
	Comment string `json:"comment,omitempty"`
	// Comment ID
	CommentId int `json:"comment_id,omitempty"`
	// Publication ID
	ContentId int `json:"content_id,omitempty"`
	// Comment date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Comment Author
	Author struct {
		// Author's reputation
		Reputation int `json:"reputation,omitempty"`
		// Author's title
		Status string `json:"status,omitempty"`
		// Title icon
		StatusImage string `json:"status_image,omitempty"`
	} `json:"author,omitempty"`
}

// WgnWargagComments Method returns comments to the content.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/comments
//
// content_id:
//     Publication ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// page_no:
//     Result page number
func (client *Client) WgnWargagComments(realm Realm, contentId int, fields []string, pageNo int) (*WgnWargagComments, error) {
	if err := ValidateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"content_id": strconv.Itoa(contentId),
		"fields": strings.Join(fields, ","),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *WgnWargagComments
	err := client.doGetDataRequest(realm, "/wgn/wargag/comments/", reqParam, &result)
	return result, err
}
