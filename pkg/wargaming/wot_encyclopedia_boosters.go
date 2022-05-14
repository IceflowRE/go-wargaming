package wargaming

import (
	"strings"
)

type WotEncyclopediaBoosters struct {
	// Personal Reserve ID
	BoosterId int `json:"booster_id,omitempty"`
	// Personal Reserve description
	Description string `json:"description,omitempty"`
	// Personal Reserve expiration time in UTC
	ExpiresAt UnixTime `json:"expires_at,omitempty"`
	// Personal Reserve auto activation flag
	IsAuto bool `json:"is_auto,omitempty"`
	// Personal Reserve duration in seconds
	Lifetime int `json:"lifetime,omitempty"`
	// Personal Reserve name
	Name string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Price in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Resource affected by Personal Reserve. Valid values: credits, experience, crew_experience, free_experience.
	Resource string `json:"resource,omitempty"`
	// Personal Reserve image
	Images struct {
		// URL to large image
		Large string `json:"large,omitempty"`
		// URL to small image
		Small string `json:"small,omitempty"`
	} `json:"images,omitempty"`
}

// WotEncyclopediaBoosters Method returns information about Personal Reserves.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/boosters
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WotEncyclopediaBoosters(realm Realm, fields []string, language string) (*WotEncyclopediaBoosters, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotEncyclopediaBoosters
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/boosters/", reqParam, &result)
	return result, err
}
