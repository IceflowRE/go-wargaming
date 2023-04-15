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
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
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
	// Images
	Images *struct {
		// Background image
		Background *string `json:"background,omitempty"`
		// Insignia image
		Insignia *string `json:"insignia,omitempty"`
	} `json:"images,omitempty"`
	// Maximum ship Tier in a season
	MaxShipTier *int `json:"max_ship_tier,omitempty"`
	// Minimum ship Tier in a season
	MinShipTier *int `json:"min_ship_tier,omitempty"`
	// Parent season ID
	ParentSeasonId *int `json:"parent_season_id,omitempty"`
	// Season ID
	SeasonId *int `json:"season_id,omitempty"`
	// Season name
	SeasonName *string `json:"season_name,omitempty"`
	// Season opening time
	StartAt *wgnTime.UnixTime `json:"start_at,omitempty"`
	// Season start rank
	StartRank *int `json:"start_rank,omitempty"`
}
