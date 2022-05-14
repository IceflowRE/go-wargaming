package wargaming

import (
	"strings"
)

type WotClansMessageboard struct {
	// Message author ID
	AuthorId int `json:"author_id,omitempty"`
	// Message creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// ID of a player who has changed the message
	EditorId int `json:"editor_id,omitempty"`
	// Indicates if the message has been read
	IsRead bool `json:"is_read,omitempty"`
	// Message text
	Message string `json:"message,omitempty"`
	// Date when message was updated
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
}

// WotClansMessageboard Method returns messages of clan message board.This method will be removed. Use method Message board (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wot/clans/messageboard
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotClansMessageboard(realm Realm, accessToken string, fields []string) (*WotClansMessageboard, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
	}

	var result *WotClansMessageboard
	err := client.doGetDataRequest(realm, "/wot/clans/messageboard/", reqParam, &result)
	return result, err
}
