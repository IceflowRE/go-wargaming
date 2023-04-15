// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClansInfoOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Extra fields that will be added to the response. Valid values:
	//
	// "private.online_members"
	Extra []string `json:"extra,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// This parameter changes members field type. Valid values:
	//
	// "id" - Members field will contain associative array with account_id indexing in response
	MembersKey *string `json:"members_key,omitempty"`
}

type ClansInfo struct {
	// Clan can invite players
	AcceptsJoinRequests *bool `json:"accepts_join_requests,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan color in HEX #RRGGBB
	Color *string `json:"color,omitempty"`
	// Clan creation date
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan creator ID
	CreatorId *int `json:"creator_id,omitempty"`
	// Clan creator's name
	CreatorName *string `json:"creator_name,omitempty"`
	// Clan description
	Description *string `json:"description,omitempty"`
	// Clan description in HTML
	DescriptionHtml *string `json:"description_html,omitempty"`
	// Information on clan emblems in games and on clan portal
	Emblems *struct {
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
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded *bool `json:"is_clan_disbanded,omitempty"`
	// Clan Commander ID
	LeaderId *int `json:"leader_id,omitempty"`
	// Commander's name
	LeaderName *string `json:"leader_name,omitempty"`
	// Information on clan members. Field format depends on members_key input parameter.
	Members *struct {
		// Player account ID
		AccountId *int `json:"account_id,omitempty"`
		// Player name
		AccountName *string `json:"account_name,omitempty"`
		// Date when player joined clan
		JoinedAt *wgnTime.UnixTime `json:"joined_at,omitempty"`
		// Technical position name
		Role *string `json:"role,omitempty"`
		// Position
		RoleI18n *string `json:"role_i18n,omitempty"`
	} `json:"members,omitempty"`
	// Number of clan members
	MembersCount *int `json:"members_count,omitempty"`
	// Clan motto
	Motto *string `json:"motto,omitempty"`
	// Clan name
	Name *string `json:"name,omitempty"`
	// Old clan name
	OldName *string `json:"old_name,omitempty"`
	// Old clan tag
	OldTag *string `json:"old_tag,omitempty"`
	// Restricted clan information
	Private *struct {
		// Clan Treasury
		ClanTreasury *struct {
			// Amount of credits in the сlan Treasury
			Credits *int `json:"credits,omitempty"`
			// Number of bonds in the сlan Treasury
			Crystal *int `json:"crystal,omitempty"`
			// Amount of gold in the сlan Treasury
			Gold *int `json:"gold,omitempty"`
		} `json:"clan_treasury,omitempty"`
		// List of clan members with an active game session in World of Tanks.
		// An extra field.
		OnlineMembers []int `json:"online_members,omitempty"`
		// Amount of gold in the сlan Treasury
		Treasury *int `json:"treasury,omitempty"`
	} `json:"private,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt *wgnTime.UnixTime `json:"renamed_at,omitempty"`
	// Clan tag
	Tag *string `json:"tag,omitempty"`
	// Time when clan details were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
