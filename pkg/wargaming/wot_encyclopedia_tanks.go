package wargaming

import (
	"strings"
)

type WotEncyclopediaTanks struct {
	// URL to outline image of vehicle
	ContourImage string `json:"contour_image,omitempty"`
	// URL to 160 x 100 px image of vehicle
	Image string `json:"image,omitempty"`
	// URL to 124 x 31 px image of vehicle
	ImageSmall string `json:"image_small,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Tier
	Level int `json:"level,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Localized name field
	NameI18N string `json:"name_i18n,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Localized nation field
	NationI18N string `json:"nation_i18n,omitempty"`
	// Localized short name of vehicle
	ShortNameI18N string `json:"short_name_i18n,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Localized vehicle type
	TypeI18N string `json:"type_i18n,omitempty"`
}

// WotEncyclopediaTanks Deprecated: Attention! The method is deprecated.
// Method returns list of all vehicles from Tankopedia.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/tanks
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
func (client *Client) WotEncyclopediaTanks(realm Realm, fields []string, language string) (*WotEncyclopediaTanks, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotEncyclopediaTanks
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/tanks/", reqParam, &result)
	return result, err
}
