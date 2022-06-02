package wargaming

import (
	"context"
	"strconv"
)

// WargagDeletecomment Method allows to delete comments. A valid access_token for the specified account is required.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/deletecomment
//
// realm:
//     Valid realms: RealmRu
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// commentId:
//     Comment ID
func (service *wgnService) WargagDeletecomment(ctx context.Context, realm Realm, accessToken string, commentId int) error {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"comment_id":   strconv.Itoa(commentId),
	}

	err := service.client.postRequest(ctx, sectionWgn, realm, "/wargag/deletecomment/", reqParam, nil)
	return err
}
