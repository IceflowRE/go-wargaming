package wargaming

import (
	"strconv"
)

type WgnWargagRate struct {
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

// WgnWargagRate Method allows to rate content. A valid access_token for the specified account is required.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/rate
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// content_id:
//     Publication ID
// rating:
//     Rate. Valid values:
//     
//     "up" &mdash; Rate content positively 
//     "down" &mdash; Rate content negatively
func (client *Client) WgnWargagRate(realm Realm, accessToken string, contentId int, rating string) (*WgnWargagRate, error) {
	if err := ValidateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"content_id": strconv.Itoa(contentId),
		"rating": rating,
	}

	var result *WgnWargagRate
	err := client.doPostDataRequest(realm, "/wgn/wargag/rate/", reqParam, &result)
	return result, err
}
