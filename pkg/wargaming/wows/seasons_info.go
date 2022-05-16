package wows

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type SeasonsInfo struct {
	// Season start rank
	StartRank int `json:"start_rank,omitempty"`
	// Season opening time
	StartAt wgnTime.UnixTime `json:"start_at,omitempty"`
	// Season name
	SeasonName string `json:"season_name,omitempty"`
	// Season ID
	SeasonId int `json:"season_id,omitempty"`
	// Parent season ID
	ParentSeasonId int `json:"parent_season_id,omitempty"`
	// Minimum ship Tier in a season
	MinShipTier int `json:"min_ship_tier,omitempty"`
	// Maximum ship Tier in a season
	MaxShipTier int `json:"max_ship_tier,omitempty"`
	// Images
	Images struct {
		// Insignia image
		Insignia string `json:"insignia,omitempty"`
		// Background image
		Background string `json:"background,omitempty"`
	} `json:"images,omitempty"`
	// Season finishing time
	FinishAt wgnTime.UnixTime `json:"finish_at,omitempty"`
	// Season closing time
	CloseAt wgnTime.UnixTime `json:"close_at,omitempty"`
	// Minimum Service Record Level to join a season
	AccountTier int `json:"account_tier,omitempty"`
} 
