package wotx

import (
	"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"
)

// ClansListOptions options.
type ClansListOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "en". Valid values:
	// 
	// "en" - English (by default)
	// "ru" - Русский 
	// "pl" - Polski 
	// "de" - Deutsch 
	// "fr" - Français 
	// "es" - Español 
	// "tr" - Türkçe
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
	// Clan emblems set ID
	EmblemSetId *int `json:"emblem_set_id,omitempty"`
	// Number of clan members
	MembersCount *int `json:"members_count,omitempty"`
	// Clan name
	Name *string `json:"name,omitempty"`
	// Clan tag
	Tag *string `json:"tag,omitempty"`
}
