// Auto generated file!

package wargaming

import (
	"context"
	"strconv"
)

// ClanmessagesLike likes or unlikes a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/like
//
// realm:
//
//	Valid realms: RealmAsia, RealmEu, RealmNa
//
// accessToken:
//
//	Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//
// action:
//
//	Action. Valid values:
//
//	"add" - Set a like for message
//	"remove" - Remove a like for message
//
// messageId:
//
//	Message ID
func (service *WotbService) ClanmessagesLike(ctx context.Context, realm Realm, accessToken string, action string, messageId int) error {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa}); err != nil {
		return err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"action":       action,
		"message_id":   strconv.Itoa(messageId),
	}

	err := service.client.postRequest(ctx, sectionWotb, realm, "/clanmessages/like/", reqParam, nil, nil)
	return err
}
