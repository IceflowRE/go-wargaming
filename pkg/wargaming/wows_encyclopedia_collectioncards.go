package wargaming

import (
	"strings"
)

type WowsEncyclopediaCollectioncards struct {
	// Item ID
	CardId int `json:"card_id,omitempty"`
	// Collection ID
	CollectionId int `json:"collection_id,omitempty"`
	// Item description
	Description string `json:"description,omitempty"`
	// Item name
	Name string `json:"name,omitempty"`
	// Item tag
	Tag string `json:"tag,omitempty"`
	// Item card icons
	Images struct {
		// URL to the large card icon
		Large string `json:"large,omitempty"`
		// URL to the medium card icon
		Medium string `json:"medium,omitempty"`
		// URL to the small card icon
		Small string `json:"small,omitempty"`
	} `json:"images,omitempty"`
}

// WowsEncyclopediaCollectioncards Method returns information about items that are included in the collection.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/collectioncards
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
func (client *Client) WowsEncyclopediaCollectioncards(realm Realm, fields []string, language string) (*WowsEncyclopediaCollectioncards, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaCollectioncards
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/collectioncards/", reqParam, &result)
	return result, err
}
