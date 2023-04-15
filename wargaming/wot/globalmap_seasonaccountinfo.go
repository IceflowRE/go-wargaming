// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type GlobalmapSeasonaccountinfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
}

type GlobalmapSeasonaccountinfo struct {
	// Account information by seasons and vehicle Tiers
	Seasons *struct {
		// Account ID
		AccountId *int `json:"account_id,omitempty"`
		// Award level
		AwardLevel *string `json:"award_level,omitempty"`
		// Battles fought for current clan
		Battles *int `json:"battles,omitempty"`
		// Battles to fight in a current clan to get clan award for the season
		BattlesToAward *int `json:"battles_to_award,omitempty"`
		// Clan ID
		ClanId *int `json:"clan_id,omitempty"`
		// Clan rating
		ClanRank *int `json:"clan_rank,omitempty"`
		// Season ID
		SeasonId *string `json:"season_id,omitempty"`
		// Date when account data were updated
		UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
		// Vehicle Tier
		VehicleLevel *int `json:"vehicle_level,omitempty"`
	} `json:"seasons,omitempty"`
}
