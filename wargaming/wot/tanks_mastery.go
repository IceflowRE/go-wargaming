// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type TanksMasteryOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Player's vehicle ID. Maximum limit: 100.
	TankId []int `json:"tank_id,omitempty"`
}

type TanksMastery struct {
	// Values of these percentiles for each piece of equipment
	Distribution map[string]map[string]int `json:"distribution,omitempty"`
	// Date of data update
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
