package wargaming

import (
	"strings"
)

type WowsClansSeason struct {
	// Number of the Rating points required for progression to the next league
	DivisionPoints int `json:"division_points,omitempty"`
	// Season end time
	FinishTime int `json:"finish_time,omitempty"`
	// Season name
	Name string `json:"name,omitempty"`
	// Season ID
	SeasonId int `json:"season_id,omitempty"`
	// Maximum ship Tier in a season
	ShipTierMax int `json:"ship_tier_max,omitempty"`
	// Minimal ship Tier in a season
	ShipTierMin int `json:"ship_tier_min,omitempty"`
	// Season start time
	StartTime int `json:"start_time,omitempty"`
	// Leagues in the season Rating
	Leagues struct {
		// Color of the league icon
		Color string `json:"color,omitempty"`
		// League icon
		Icon string `json:"icon,omitempty"`
		// League name
		Name string `json:"name,omitempty"`
	} `json:"leagues,omitempty"`
}

// WowsClansSeason Method returns information about Clan Battles season.
//
// https://developers.wargaming.net/reference/all/wows/clans/season
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
func (client *Client) WowsClansSeason(realm Realm, fields []string, language string) (*WowsClansSeason, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsClansSeason
	err := client.doGetDataRequest(realm, "/wows/clans/season/", reqParam, &result)
	return result, err
}
