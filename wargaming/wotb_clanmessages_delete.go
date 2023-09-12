// Auto generated file!

package wargaming

import (
	"context"
	"strconv"
)

// ClanmessagesDelete removes a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/delete
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
//
// accessToken:
//
//	Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//
// messageId:
//
//	Message ID
func (service *WotbService) ClanmessagesDelete(ctx context.Context, realm Realm, accessToken string, messageId int) error {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"message_id":   strconv.Itoa(messageId),
	}

	err := service.client.postRequest(ctx, sectionWotb, realm, "/clanmessages/delete/", reqParam, nil, nil)
	return err
}
