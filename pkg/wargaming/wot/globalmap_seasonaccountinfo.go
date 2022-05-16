package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapSeasonaccountinfo struct {
	// Account information by seasons and vehicle Tiers
	Seasons struct {
		// Vehicle Tier
		VehicleLevel int `json:"vehicle_level,omitempty"`
		// Date when account data were updated
		UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
		// Season ID
		SeasonId string `json:"season_id,omitempty"`
		// Clan rating
		ClanRank int `json:"clan_rank,omitempty"`
		// Clan ID
		ClanId int `json:"clan_id,omitempty"`
		// Battles to fight in a current clan to get clan award for the season
		BattlesToAward int `json:"battles_to_award,omitempty"`
		// Battles fought for current clan
		Battles int `json:"battles,omitempty"`
		// Award level
		AwardLevel string `json:"award_level,omitempty"`
		// Account ID
		AccountId int `json:"account_id,omitempty"`
	} `json:"seasons,omitempty"`
} 
