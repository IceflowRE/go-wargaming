package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WgnClansMembersinfo struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Player name
	AccountName string `json:"account_name,omitempty"`
	// Date when player joined clan
	JoinedAt UnixTime `json:"joined_at,omitempty"`
	// Technical position name
	Role string `json:"role,omitempty"`
	// Position
	RoleI18N string `json:"role_i18n,omitempty"`
	// Short clan info
	Clan struct {
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
		// Clan color in HEX #RRGGBB
		Color string `json:"color,omitempty"`
		// Clan creation date
		CreatedAt UnixTime `json:"created_at,omitempty"`
		// Game where clan was created
		Game string `json:"game,omitempty"`
		// Number of clan members
		MembersCount int `json:"members_count,omitempty"`
		// Clan name
		Name string `json:"name,omitempty"`
		// Clan tag
		Tag string `json:"tag,omitempty"`
		// Information on clan emblems in games and on clan portal
		Emblems struct {
			// List of links to 195x195 px emblems
			X195 map[string]string `json:"x195,omitempty"`
			// List of links to 24x24 px emblems
			X24 map[string]string `json:"x24,omitempty"`
			// List of links to 256x256 px emblems
			X256 map[string]string `json:"x256,omitempty"`
			// List of links to 32x32 px emblems
			X32 map[string]string `json:"x32,omitempty"`
			// List of links to 64x64 px emblems
			X64 map[string]string `json:"x64,omitempty"`
		} `json:"emblems,omitempty"`
	} `json:"clan,omitempty"`
}

// WgnClansMembersinfo Deprecated: Attention! The method is deprecated.
// Method returns detailed clan member information and brief clan details. Information is available for World of Tanks and World of Warplanes clans.This method will be removed. Use method Clan member details (World of Tanks)
//
// https://developers.wargaming.net/reference/all/wgn/clans/membersinfo
//
// account_id:
//     Account ID. Maximum limit: 100. Min value is 1.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// game:
//     Name of the game to perform the clan search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Valid values:
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
func (client *Client) WgnClansMembersinfo(realm Realm, accountId []int, fields []string, game string, language string) (*WgnClansMembersinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"fields": strings.Join(fields, ","),
		"game": game,
		"language": language,
	}

	var result *WgnClansMembersinfo
	err := client.doGetDataRequest(realm, "/wgn/clans/membersinfo/", reqParam, &result)
	return result, err
}
