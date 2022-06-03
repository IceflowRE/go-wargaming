package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
)

// WargagNewcomment Method allows to add comments. A valid access_token for the specified account is required. Comments allowed for picture and video content only.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/newcomment
//
// realm:
//     Valid realms: RealmRu
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// contentId:
//     Publication ID
// text:
//     Comment text. Maximum length: 1000 symbols. Maximum length: 1000.
func (service *wgnService) WargagNewcomment(ctx context.Context, realm Realm, accessToken string, contentId int, text string) (*wgn.WargagNewcomment, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"content_id":   strconv.Itoa(contentId),
		"text":         text,
	}

	var data *wgn.WargagNewcomment
	err := service.client.postRequest(ctx, sectionWgn, realm, "/wargag/newcomment/", reqParam, &data)
	return data, err
}
