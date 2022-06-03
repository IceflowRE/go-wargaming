package wot

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// ClansListOptions options.
type ClansListOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	// 
	// "en" - English 
	// "ru" - Русский (by default)
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
	Language *string
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int
	// Page number. Default is 1. Min value is 1.
	PageNo *int
	// Part of name or tag for clan search. Minimum 2 characters
	Search *string
}

type ClansList struct {
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
}
