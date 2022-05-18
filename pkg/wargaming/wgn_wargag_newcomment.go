package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
)

// WgnWargagNewcomment Method allows to add comments. A valid access_token for the specified account is required. Comments allowed for picture and video content only.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/newcomment
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// contentId:
//     Publication ID
// text:
//     Comment text. Maximum length: 1000 symbols. Maximum length: 1000.
func (client *Client) WgnWargagNewcomment(realm Realm, accessToken string, contentId int, text string) (*wgn.WargagNewcomment, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"content_id": strconv.Itoa(contentId),
		"text": text,
	}

	var result *wgn.WargagNewcomment
	err := client.doPostDataRequest(realm, "/wgn/wargag/newcomment/", reqParam, &result)
	return result, err
}
