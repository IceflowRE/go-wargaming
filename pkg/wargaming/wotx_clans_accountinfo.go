package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotxClansAccountinfo struct {
	// User ID
	AccountId int `json:"account_id,omitempty"`
	// Player name
	AccountName string `json:"account_name,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Cooldown for leaving the clan
	InClanCooldownTill UnixTime `json:"in_clan_cooldown_till,omitempty"`
	// Date when player joined clan
	JoinedAt UnixTime `json:"joined_at,omitempty"`
	// Technical position name
	Role string `json:"role,omitempty"`
	// Brief clan details.
	// An extra field.
	Clan struct {
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
		// Clan color in HEX #RRGGBB
		Color string `json:"color,omitempty"`
		// Clan creation date
		CreatedAt UnixTime `json:"created_at,omitempty"`
		// Emblems set ID
		EmblemSetId int `json:"emblem_set_id,omitempty"`
		// Number of clan members
		MembersCount int `json:"members_count,omitempty"`
		// Clan name
		Name string `json:"name,omitempty"`
		// Clan tag
		Tag string `json:"tag,omitempty"`
	} `json:"clan,omitempty"`
}

// WotxClansAccountinfo Method returns player clan data. Player clan data exist only for accounts, that were clan members at least once.
//
// https://developers.wargaming.net/reference/all/wotx/clans/accountinfo
//
// account_id:
//     Account ID. Maximum limit: 100. Min value is 1.
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "clan"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
func (client *Client) WotxClansAccountinfo(realm Realm, accountId []int, extra []string, fields []string, language string) (*WotxClansAccountinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxClansAccountinfo
	err := client.doGetDataRequest(realm, "/wotx/clans/accountinfo/", reqParam, &result)
	return result, err
}
