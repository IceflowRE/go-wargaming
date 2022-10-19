// Auto generated file!

package wows

type EncyclopediaCollectionsOptions struct {
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

type EncyclopediaCollections struct {
	// Collection item cost in duplicates
	CardCost *int `json:"card_cost,omitempty"`
	// Collection ID
	CollectionId *int `json:"collection_id,omitempty"`
	// Collection description
	Description *string `json:"description,omitempty"`
	// Link to collection image
	Image *string `json:"image,omitempty"`
	// Collection name
	Name *string `json:"name,omitempty"`
	// Collection tag
	Tag *string `json:"tag,omitempty"`
}
