// Auto generated file!

package wows

type SeasonsShipstatsOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
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
	// Ship ID. Maximum limit: 100.
	ShipId []int `json:"ship_id,omitempty"`
}

type SeasonsShipstats struct {
	// User ID
	AccountId *int `json:"account_id,omitempty"`
	// Ranked Battles seasons
	Seasons map[string]string `json:"seasons,omitempty"`
	// Ship ID
	ShipId *int `json:"ship_id,omitempty"`
}
