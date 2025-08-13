// Auto generated file!

package wows

type ShipsBadgesOptions struct {
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
	// Player's ship ID. Maximum limit: 100.
	ShipId []int `json:"ship_id,omitempty"`
}

type ShipsBadges struct {
	// Ship ID
	ShipId *int `json:"ship_id,omitempty"`
	// Top grade mastery badge
	TopGradeClass *int `json:"top_grade_class,omitempty"`
}
