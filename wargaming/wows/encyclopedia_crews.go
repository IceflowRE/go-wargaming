// Auto generated file!

package wows

type EncyclopediaCrewsOptions struct {
	// Commander ID. Maximum limit: 100.
	CommanderId []int `json:"commander_id,omitempty"`
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

type EncyclopediaCrews struct {
	// Basic training cost
	BaseTrainingHirePrice *int `json:"base_training_hire_price,omitempty"`
	// Basic training level
	BaseTrainingLevel *int `json:"base_training_level,omitempty"`
	// Commanders' first names
	FirstNames []string `json:"first_names,omitempty"`
	// Retraining cost in doubloons
	GoldRetrainingPrice *int `json:"gold_retraining_price,omitempty"`
	// Training cost in gold
	GoldTrainingHirePrice *int `json:"gold_training_hire_price,omitempty"`
	// Commander training level purchased for doubloons
	GoldTrainingLevel *int `json:"gold_training_level,omitempty"`
	// The list of the Commander images:
	//
	// 1-URL to the image of the Commander with 1–7 skill points;
	// 8-URL to the image of the Commander with 8–13 skill points;
	// 14-URL to the image of the Commander with 14–20 skill points;.
	//
	// If only the value for the key 1 is specified, the Commander has not earned skill points yet.
	Icons map[string]string `json:"icons,omitempty"`
	// Indicates if the skill is preserved after retraining
	IsRetrainable *bool `json:"is_retrainable,omitempty"`
	// Commanders' last names
	LastNames []string `json:"last_names,omitempty"`
	// Retraining cost in credits
	MoneyRetrainingPrice *int `json:"money_retraining_price,omitempty"`
	// Training cost in credits
	MoneyTrainingHirePrice *int `json:"money_training_hire_price,omitempty"`
	// Commander training level purchased for credits
	MoneyTrainingLevel *int `json:"money_training_level,omitempty"`
	// Nation
	Nation *string `json:"nation,omitempty"`
	// Subnation index
	SubnationIndex *int `json:"subnation_index,omitempty"`
}
