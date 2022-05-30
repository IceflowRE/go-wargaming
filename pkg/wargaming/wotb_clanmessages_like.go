package wargaming

import (
	"strconv"
)

// WotbClanmessagesLike Method likes or unlikes a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/like
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//     Parameter is required.
// action:
//     Action. Valid values:
//     
//     "add" - Set a like for message 
//     "remove" - Remove a like for message
//     Parameter is required.
// messageId:
//     Message ID
//     Parameter is required.
func (client *Client) WotbClanmessagesLike(realm Realm, accessToken string, action string, messageId int) error {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"action": action,
		"message_id": strconv.Itoa(messageId),
	}

	err := client.doPostRequest(realm, "/wotb/clanmessages/like/", reqParam)
	return err
}
