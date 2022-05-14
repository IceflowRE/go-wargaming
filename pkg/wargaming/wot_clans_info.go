package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotClansInfo struct {
	// Clan can invite players
	AcceptsJoinRequests bool `json:"accepts_join_requests,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan creation date
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan description
	Description string `json:"description,omitempty"`
	// Clan description in HTML
	DescriptionHtml string `json:"description_html,omitempty"`
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded bool `json:"is_clan_disbanded,omitempty"`
	// Clan Commander ID
	LeaderId int `json:"leader_id,omitempty"`
	// Commander's name
	LeaderName string `json:"leader_name,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Clan motto
	Motto string `json:"motto,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Old clan name
	OldName string `json:"old_name,omitempty"`
	// Old clan tag
	OldTag string `json:"old_tag,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt UnixTime `json:"renamed_at,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Time when clan details were updated
	UpdatedAt UnixTime `json:"updated_at,omitempty"`
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
	// Information on clan members. Field format depends on members_key input parameter.
	Members struct {
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
	} `json:"members,omitempty"`
	// Restricted clan information
	Private struct {
		// List of clan members with an active game session in World of Tanks.
		// An extra field.
		OnlineMembers []int `json:"online_members,omitempty"`
		// Amount of gold in the сlan Treasury
		Treasury int `json:"treasury,omitempty"`
		// Clan Treasury
		ClanTreasury struct {
			// Amount of credits in the сlan Treasury
			Credits int `json:"credits,omitempty"`
			// Number of bonds in the сlan Treasury
			Crystal int `json:"crystal,omitempty"`
			// Amount of gold in the сlan Treasury
			Gold int `json:"gold,omitempty"`
		} `json:"clan_treasury,omitempty"`
	} `json:"private,omitempty"`
}

// WotClansInfo Method returns detailed clan information.
//
// https://developers.wargaming.net/reference/all/wot/clans/info
//
// clan_id:
//     Clan ID. Maximum limit: 100.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// extra:
//     Extra fields that will be added to the response. Valid values:
//     
//     "private.online_members"
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
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
// members_key:
//     This parameter changes members field type. Valid values:
//     
//     "id" &mdash; Members field will contain associative array with account_id indexing in response
func (client *Client) WotClansInfo(realm Realm, clanId []int, accessToken string, extra []string, fields []string, language string, membersKey string) (*WotClansInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"clan_id": utils.SliceIntToString(clanId, ","),
		"access_token": accessToken,
		"extra": strings.Join(extra, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"members_key": membersKey,
	}

	var result *WotClansInfo
	err := client.doGetDataRequest(realm, "/wot/clans/info/", reqParam, &result)
	return result, err
}
