package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strconv"
	"strings"
)

// WotbClanmessagesMessages Method returns messages of clan message board. If a single message requested by ID, all filters are ignored and total equals null.
//
// https://developers.wargaming.net/reference/all/wotb/clanmessages/messages
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// expiresAfter:
//     Search for messages with end date of relevance equal or after the value. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// expiresBefore:
//     Search for messages with end date of relevance before the value. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// importance:
//     Message importance. Valid values:
//     
//     "important" - Important messages 
//     "standard" - Standard messages
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" - English 
//     "ru" - Русский (by default)
//     "pl" - Polski 
//     "de" - Deutsch 
//     "fr" - Français 
//     "es" - Español 
//     "zh-cn" - 简体中文 
//     "zh-tw" - 繁體中文 
//     "tr" - Türkçe 
//     "cs" - Čeština 
//     "th" - ไทย 
//     "vi" - Tiếng Việt 
//     "ko" - 한국어
// limit:
//     Number of returned entries. If the value sent exceeds 100, a limit of 25 is applied (by default). Default is 25. Min value is 1.
// messageId:
//     Message ID
// orderBy:
//     Sorting. Valid values:
//     
//     
//     importance — by message importance
//     
//     
//     -importance — by message importance in reverse order
//     
//     
//     created_at — by date of message creation
//     
//     
//     -created_at — by date of message creation in reverse order
//     
//     
//     updated_at — by message update date
//     
//     
//     -updated_at — by message update date in reverse order
//     
//     
//     type — by message type
//     
//     
//     -type — by message type in reverse order. Default is "importance, type". Maximum limit: 100.
// pageNo:
//     Page number. Default is 1. Min value is 1.
// status:
//     Message status. Valid values:
//     
//     "active" - Active message 
//     "deleted" - Deleted message
// type_:
//     Message type. Valid values:
//     
//     "general" - General messages 
//     "training" - Training messages 
//     "meeting" - Meeting messages 
//     "battle" - Battle messages
func (client *Client) WotbClanmessagesMessages(realm Realm, accessToken string, expiresAfter wgnTime.UnixTime, expiresBefore wgnTime.UnixTime, fields []string, importance string, language string, limit int, messageId int, orderBy []string, pageNo int, status string, type_ string) (*wotb.ClanmessagesMessages, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"expires_after": strconv.FormatInt(expiresAfter.Unix(), 10),
		"expires_before": strconv.FormatInt(expiresBefore.Unix(), 10),
		"fields": strings.Join(fields, ","),
		"importance": importance,
		"language": language,
		"limit": strconv.Itoa(limit),
		"message_id": strconv.Itoa(messageId),
		"order_by": strings.Join(orderBy, ","),
		"page_no": strconv.Itoa(pageNo),
		"status": status,
		"type": type_,
	}

	var result *wotb.ClanmessagesMessages
	err := client.doGetDataRequest(realm, "/wotb/clanmessages/messages/", reqParam, &result)
	return result, err
}
