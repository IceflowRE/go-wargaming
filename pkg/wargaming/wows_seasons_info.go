package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsSeasonsInfo struct {
	// Minimum Service Record Level to join a season
	AccountTier int `json:"account_tier,omitempty"`
	// Season closing time
	CloseAt UnixTime `json:"close_at,omitempty"`
	// Season finishing time
	FinishAt UnixTime `json:"finish_at,omitempty"`
	// Maximum ship Tier in a season
	MaxShipTier int `json:"max_ship_tier,omitempty"`
	// Minimum ship Tier in a season
	MinShipTier int `json:"min_ship_tier,omitempty"`
	// Parent season ID
	ParentSeasonId int `json:"parent_season_id,omitempty"`
	// Season ID
	SeasonId int `json:"season_id,omitempty"`
	// Season name
	SeasonName string `json:"season_name,omitempty"`
	// Season opening time
	StartAt UnixTime `json:"start_at,omitempty"`
	// Season start rank
	StartRank int `json:"start_rank,omitempty"`
	// Images
	Images struct {
		// Background image
		Background string `json:"background,omitempty"`
		// Insignia image
		Insignia string `json:"insignia,omitempty"`
	} `json:"images,omitempty"`
}

// WowsSeasonsInfo Method returns information about Ranked Battles seasons.
//
// https://developers.wargaming.net/reference/all/wows/seasons/info
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
// season_id:
//     Season ID. Maximum limit: 100.
func (client *Client) WowsSeasonsInfo(realm Realm, fields []string, language string, seasonId []int) (*WowsSeasonsInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"season_id": utils.SliceIntToString(seasonId, ","),
	}

	var result *WowsSeasonsInfo
	err := client.doGetDataRequest(realm, "/wows/seasons/info/", reqParam, &result)
	return result, err
}
