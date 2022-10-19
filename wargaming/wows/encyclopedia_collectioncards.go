// Auto generated file!

package wows

type EncyclopediaCollectioncardsOptions struct {
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

type EncyclopediaCollectioncards struct {
	// Item ID
	CardId *int `json:"card_id,omitempty"`
	// Collection ID
	CollectionId *int `json:"collection_id,omitempty"`
	// Item description
	Description *string `json:"description,omitempty"`
	// Item card icons
	Images *struct {
		// URL to the large card icon
		Large *string `json:"large,omitempty"`
		// URL to the medium card icon
		Medium *string `json:"medium,omitempty"`
		// URL to the small card icon
		Small *string `json:"small,omitempty"`
	} `json:"images,omitempty"`
	// Item name
	Name *string `json:"name,omitempty"`
	// Item tag
	Tag *string `json:"tag,omitempty"`
}
