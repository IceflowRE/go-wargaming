// Auto generated file!

package wows

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type SeasonsInfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "it" - Italiano
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "zh-cn" - 中文
	// "tr" - Türkçe
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
	// Season ID. Maximum limit: 100.
	SeasonId []int `json:"season_id,omitempty"`
}

type SeasonsInfo struct {
	// Minimum Service Record Level to join a season
	AccountTier *int `json:"account_tier,omitempty"`
	// Season closing time
	CloseAt *wgnTime.UnixTime `json:"close_at,omitempty"`
	// Season finishing time
	FinishAt *wgnTime.UnixTime `json:"finish_at,omitempty"`
	// Leagues
	Leagues *struct {
		// Images
		Images *struct {
			// Inactive image of League
			Inactive *string `json:"inactive,omitempty"`
			// Normal image of League
			Normal *string `json:"normal,omitempty"`
		} `json:"images,omitempty"`
		// League Name
		LeagueName *string `json:"league_name,omitempty"`
		// Maximum ship Tier in a league
		MaxShipTier *int `json:"max_ship_tier,omitempty"`
		// Minimum ship Tier in a league
		MinShipTier *int `json:"min_ship_tier,omitempty"`
		// League ranks
		Ranks *struct {
			// Images
			Images *struct {
				// Default image of league ranks
				Default *string `json:"default,omitempty"`
				// Small image of league ranks
				Small *string `json:"small,omitempty"`
			} `json:"images,omitempty"`
			// Save point
			IsSavePoint *bool `json:"is_save_point,omitempty"`
			// Number of stars to lose the current rank
			StarLossPlace *int `json:"star_loss_place,omitempty"`
			// Number of stars to get the next rank
			StarsToNext *int `json:"stars_to_next,omitempty"`
		} `json:"ranks,omitempty"`
		// Limit of ships in a league
		ShipTotalLimit *int `json:"ship_total_limit,omitempty"`
		// League start rank
		StartRank *int `json:"start_rank,omitempty"`
	} `json:"leagues,omitempty"`
	// Season ID
	SeasonId *int `json:"season_id,omitempty"`
	// Season name
	SeasonName *string `json:"season_name,omitempty"`
	// Season opening time
	StartAt *wgnTime.UnixTime `json:"start_at,omitempty"`
}
