package wargaming

import (
	"context"
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgn"
	"strconv"
)

// WargagRate allows to rate content. A valid access_token for the specified account is required.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/rate
//
// realm:
//     Valid realms: RealmRu
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// contentId:
//     Publication ID
// rating:
//     Rate. Valid values:
//
//     "up" - Rate content positively
//     "down" - Rate content negatively
func (service *WgnService) WargagRate(ctx context.Context, realm Realm, accessToken string, contentId int, rating string) (*wgn.WargagRate, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"content_id":   strconv.Itoa(contentId),
		"rating":       rating,
	}

	var data *wgn.WargagRate
	err := service.client.postRequest(ctx, sectionWgn, realm, "/wargag/rate/", reqParam, &data)
	return data, err
}
