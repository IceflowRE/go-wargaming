package wargaming

import (
	"strconv"
)

// WotbClanmessagesLike Method likes or unlikes a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/like
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// action:
//     Action. Valid values:
//     
//     "add" &mdash; Set a like for message 
//     "remove" &mdash; Remove a like for message
// message_id:
//     Message ID
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
