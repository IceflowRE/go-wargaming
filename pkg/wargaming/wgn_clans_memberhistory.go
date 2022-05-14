package wargaming

import (
	"strings"
	"strconv"
)

type WgnClansMemberhistory struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Date when player joined clan
	JoinedAt UnixTime `json:"joined_at,omitempty"`
	// Date when player left clan
	LeftAt UnixTime `json:"left_at,omitempty"`
	// Last position in clan
	Role string `json:"role,omitempty"`
}

// WgnClansMemberhistory Deprecated: Attention! The method is deprecated.
// Method returns information about player's clan history. Data on 10 last clan memberships are presented in the response.This method will be removed. Use method Player's clan history (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/memberhistory
//
// account_id:
//     Account ID. Min value is 1.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to perform the clan history search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Valid values:
//     
//     "wot" &mdash; World of Tanks (by default)
//     "wowp" &mdash; World of Warplanes
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WgnClansMemberhistory(realm Realm, accountId int, fields []string, game string, language string) (*WgnClansMemberhistory, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"fields": strings.Join(fields, ","),
		"game": game,
		"language": language,
	}

	var result *WgnClansMemberhistory
	err := client.doGetDataRequest(realm, "/wgn/clans/memberhistory/", reqParam, &result)
	return result, err
}
