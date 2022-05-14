package wargaming

import (
	"strings"
	"strconv"
)

type WgnWargagSearch struct {
	// Publication author ID
	AccountId int `json:"account_id,omitempty"`
	// Content category ID
	CategoryId int `json:"category_id,omitempty"`
	// Publication ID
	ContentId int `json:"content_id,omitempty"`
	// Publication date in UNIX timestamp or ISO 8601
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Publication text
	Description string `json:"description,omitempty"`
	// Link to image preview. Available for picture content only.
	MediaPreviewUrl string `json:"media_preview_url,omitempty"`
	// Video/image link
	MediaUrl string `json:"media_url,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Current rating
	Rating int `json:"rating,omitempty"`
	// Publication topic
	Subject string `json:"subject,omitempty"`
	// Tag ID
	TagId int `json:"tag_id,omitempty"`
	// Content type
	Type_ string `json:"type,omitempty"`
	// Link to original publication
	WargagUrl string `json:"wargag_url,omitempty"`
	// Indicates the possibility to vote for content. This data requires a valid access_token for the specified account.
	AllowedToVote bool `json:"allowed_to_vote,omitempty"`
	// Content author
	Author struct {
		// Author's reputation
		Reputation int `json:"reputation,omitempty"`
		// Author's title
		Status string `json:"status,omitempty"`
		// Title icon
		StatusImage string `json:"status_image,omitempty"`
	} `json:"author,omitempty"`
}

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
func (client *Client) WgnWargagSearch(realm Realm, q string, accessToken string, categoryId int, fields []string, ratingThreshold int, tagId int, type_ string) (*WgnWargagSearch, error) {
	if err := ValidateRealm(realm, []Realm{RealmRu}); err != nil {
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

	var result *WgnWargagSearch
	err := client.doGetDataRequest(realm, "/wgn/wargag/search/", reqParam, &result)
	return result, err
}
