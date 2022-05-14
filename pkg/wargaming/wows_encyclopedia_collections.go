package wargaming

import (
	"strings"
)

type WowsEncyclopediaCollections struct {
	// Collection item cost in duplicates
	CardCost int `json:"card_cost,omitempty"`
	// Collection ID
	CollectionId int `json:"collection_id,omitempty"`
	// Collection description
	Description string `json:"description,omitempty"`
	// Link to collection image
	Image string `json:"image,omitempty"`
	// Collection name
	Name string `json:"name,omitempty"`
	// Collection tag
	Tag string `json:"tag,omitempty"`
}

// WowsEncyclopediaCollections Method returns information about collections.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/collections
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
func (client *Client) WowsEncyclopediaCollections(realm Realm, fields []string, language string) (*WowsEncyclopediaCollections, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsEncyclopediaCollections
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/collections/", reqParam, &result)
	return result, err
}
