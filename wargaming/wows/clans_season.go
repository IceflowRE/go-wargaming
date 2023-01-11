// Auto generated file!

package wows

type ClansSeasonOptions struct {
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
}

type ClansSeason struct {
	// Number of the Rating points required for progression to the next league
	DivisionPoints *int `json:"division_points,omitempty"`
	// Season end time
	FinishTime *int `json:"finish_time,omitempty"`
	// Leagues in the season Rating
	Leagues *struct {
		// Color of the league icon
		Color *string `json:"color,omitempty"`
		// League icon
		Icon *string `json:"icon,omitempty"`
		// League name
		Name *string `json:"name,omitempty"`
	} `json:"leagues,omitempty"`
	// Season name
	Name *string `json:"name,omitempty"`
	// Season ID
	SeasonId *int `json:"season_id,omitempty"`
	// Maximum ship Tier in a season
	ShipTierMax *int `json:"ship_tier_max,omitempty"`
	// Minimal ship Tier in a season
	ShipTierMin *int `json:"ship_tier_min,omitempty"`
	// Season start time
	StartTime *int `json:"start_time,omitempty"`
}
