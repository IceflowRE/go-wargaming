// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClansAccountinfoOptions struct {
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
}

type ClansAccountinfo struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Player name
	AccountName *string `json:"account_name,omitempty"`
	// Short clan info
	Clan *struct {
		// Clan ID
		ClanId *int `json:"clan_id,omitempty"`
		// Clan color in HEX #RRGGBB
		Color *string `json:"color,omitempty"`
		// Clan creation date
		CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
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
		// Number of clan members
		MembersCount *int `json:"members_count,omitempty"`
		// Clan name
		Name *string `json:"name,omitempty"`
		// Clan tag
		Tag *string `json:"tag,omitempty"`
	} `json:"clan,omitempty"`
	// Date when player joined clan
	JoinedAt *wgnTime.UnixTime `json:"joined_at,omitempty"`
	// Technical position name
	Role *string `json:"role,omitempty"`
	// Position
	RoleI18n *string `json:"role_i18n,omitempty"`
}
