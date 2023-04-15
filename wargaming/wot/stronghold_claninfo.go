// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type StrongholdClaninfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "pl" - Polish
	// "fr" - French
	// "es" - Spanish
	// "cs" - Czech
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
}

type StrongholdClaninfo struct {
	// Statistics for the clan's battles in the Stronghold mode
	BattlesForStrongholdsStatistics *struct {
		// End time of the last battle held on Tier X vehicles
		LastTime10 *wgnTime.UnixTime `json:"last_time_10,omitempty"`
		// Number of defeats on Tier X vehicles
		Lose10 *int `json:"lose_10,omitempty"`
		// Total number of battles on Tier X vehicles
		Total10 *int `json:"total_10,omitempty"`
		// Total number of battles on Tier X vehicles within the last 28 days
		Total10In28d *int `json:"total_10_in_28d,omitempty"`
		// Number of victories on Tier X vehicles
		Win10 *int `json:"win_10,omitempty"`
		// Number of victories on Tier X vehicles within the last 28 days
		Win10In28d *int `json:"win_10_in_28d,omitempty"`
	} `json:"battles_for_strongholds_statistics,omitempty"`
	// Statistics for skirmishes against the clan's Stronghold
	BattlesSeriesForStrongholdsStatistics *struct {
		// Number of defeats on Tier X vehicles
		Lose10 *int `json:"lose_10,omitempty"`
		// Total number of battles on Tier X vehicles
		Total10 *int `json:"total_10,omitempty"`
		// Total number of battles on Tier X vehicles within the last 28 days
		Total10In28d *int `json:"total_10_in_28d,omitempty"`
		// Number of victories on Tier X vehicles
		Win10 *int `json:"win_10,omitempty"`
		// Number of victories on Tier X vehicles within the last 28 days
		Win10In28d *int `json:"win_10_in_28d,omitempty"`
	} `json:"battles_series_for_strongholds_statistics,omitempty"`
	// Information about the Stronghold's construction sites
	BuildingSlots *struct {
		// Map ID associated with the current construction site
		ArenaId *string `json:"arena_id,omitempty"`
		// Level of the structure located on the current construction site
		BuildingLevel *int `json:"building_level,omitempty"`
		// Name of the structure located on the current construction site
		BuildingTitle *string `json:"building_title,omitempty"`
		// Bridgehead's name
		Direction *string `json:"direction,omitempty"`
		// Position of the construction site
		Position *string `json:"position,omitempty"`
		// Name of the Reserve earned in the structure located on the current construction site
		ReserveTitle *string `json:"reserve_title,omitempty"`
	} `json:"building_slots,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan name
	ClanName *string `json:"clan_name,omitempty"`
	// Clan tag
	ClanTag *string `json:"clan_tag,omitempty"`
	// Map ID associated with the Command Center construction site
	CommandCenterArenaId *string `json:"command_center_arena_id,omitempty"`
	// Clan's skirmish statistics
	SkirmishStatistics *struct {
		// End time of the last battle held on Tier X vehicles
		LastTime10 *wgnTime.UnixTime `json:"last_time_10,omitempty"`
		// End time of the last battle held on Tier VI vehicles
		LastTime6 *wgnTime.UnixTime `json:"last_time_6,omitempty"`
		// End time of the last battle held on Tier VIII vehicles
		LastTime8 *wgnTime.UnixTime `json:"last_time_8,omitempty"`
		// Number of defeats on Tier X vehicles
		Lose10 *int `json:"lose_10,omitempty"`
		// Number of defeats on Tier VI vehicles
		Lose6 *int `json:"lose_6,omitempty"`
		// Number of defeats on Tier VIII vehicles
		Lose8 *int `json:"lose_8,omitempty"`
		// Total number of battles on Tier X vehicles
		Total10 *int `json:"total_10,omitempty"`
		// Total number of battles on Tier X vehicles within the last 28 days
		Total10In28d *int `json:"total_10_in_28d,omitempty"`
		// Total number of battles on Tier VI vehicles
		Total6 *int `json:"total_6,omitempty"`
		// Total number of battles on Tier VI vehicles within the last 28 days
		Total6In28d *int `json:"total_6_in_28d,omitempty"`
		// Total number of battles on Tier VIII vehicles
		Total8 *int `json:"total_8,omitempty"`
		// Total number of battles on Tier VIII vehicles within the last 28 days
		Total8In28d *int `json:"total_8_in_28d,omitempty"`
		// Number of victories on Tier X vehicles
		Win10 *int `json:"win_10,omitempty"`
		// Number of victories on Tier X vehicles within the last 28 days
		Win10In28d *int `json:"win_10_in_28d,omitempty"`
		// Number of victories on Tier VI vehicles
		Win6 *int `json:"win_6,omitempty"`
		// Number of victories on Tier VI vehicles within the last 28 days
		Win6In28d *int `json:"win_6_in_28d,omitempty"`
		// Number of victories on Tier VIII vehicles
		Win8 *int `json:"win_8,omitempty"`
		// Number of victories on Tier VIII vehicles within the last 28 days
		Win8In28d *int `json:"win_8_in_28d,omitempty"`
	} `json:"skirmish_statistics,omitempty"`
	// Total level of the Stronghold's structures
	StrongholdBuildingsLevel *int `json:"stronghold_buildings_level,omitempty"`
	// Stronghold's level
	StrongholdLevel *int `json:"stronghold_level,omitempty"`
}
