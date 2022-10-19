// Auto generated file!

package wot

type TanksAchievementsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Filter by vehicle availability in the Garage. If the parameter is not specified, all vehicles are returned. Parameter processing requires a valid access_token for the specified account_id. Valid values:
	//
	// "1" - Return vehicles available in the Garage.
	// "0" - Return vehicles that are no longer in the Garage.
	InGarage *string `json:"in_garage,omitempty"`
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
	// Player's vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
}

type TanksAchievements struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Current values of Achievement Series
	Series map[string]string `json:"series,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
}
