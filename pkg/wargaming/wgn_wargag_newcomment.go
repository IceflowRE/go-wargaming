package wargaming

import (
	"strconv"
)

type WgnWargagNewcomment struct {
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

// WgnWargagNewcomment Method allows to add comments. A valid access_token for the specified account is required. Comments allowed for picture and video content only.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/newcomment
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// content_id:
//     Publication ID
// text:
//     Comment text. Maximum length: 1000 symbols. Maximum length: 1000.
func (client *Client) WgnWargagNewcomment(realm Realm, accessToken string, contentId int, text string) (*WgnWargagNewcomment, error) {
	if err := ValidateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"content_id": strconv.Itoa(contentId),
		"text": text,
	}

	var result *WgnWargagNewcomment
	err := client.doPostDataRequest(realm, "/wgn/wargag/newcomment/", reqParam, &result)
	return result, err
}
