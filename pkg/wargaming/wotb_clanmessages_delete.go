package wargaming

import (
	"strconv"
)

// WotbClanmessagesDelete Method removes a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/delete
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// message_id:
//     Message ID
func (client *Client) WotbClanmessagesDelete(realm Realm, accessToken string, messageId int) error {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"message_id": strconv.Itoa(messageId),
	}

	err := client.doPostRequest(realm, "/wotb/clanmessages/delete/", reqParam)
	return err
}
