package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
)

// WotbClanmessagesCreate Method creates a message on clan message board.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/create
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// expiresAt:
//     Date when message will become irrelevant. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00. Date must be in the future.
// importance:
//     Message importance. Valid values:
//     
//     "important" &mdash; Important messages 
//     "standard" &mdash; Standard messages
// text:
//     Message text. Max. length: 1000. Maximum length: 1000.
// title:
//     Message title. Max. length: 100. Maximum length: 100.
// type_:
//     Message type. Valid values:
//     
//     "general" &mdash; General messages 
//     "training" &mdash; Training messages 
//     "meeting" &mdash; Meeting messages 
//     "battle" &mdash; Battle messages
func (client *Client) WotbClanmessagesCreate(realm Realm, accessToken string, expiresAt wgnTime.UnixTime, importance string, text string, title string, type_ string) (*wotb.ClanmessagesCreate, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"expires_at": strconv.FormatInt(expiresAt.Unix(), 10),
		"importance": importance,
		"text": text,
		"title": title,
		"type": type_,
	}

	var result *wotb.ClanmessagesCreate
	err := client.doPostDataRequest(realm, "/wotb/clanmessages/create/", reqParam, &result)
	return result, err
}
