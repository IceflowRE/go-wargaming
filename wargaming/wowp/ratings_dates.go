// Auto generated file!

package wowp

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type RatingsDatesOptions struct {
	// Player account ID. Maximum limit: 100.
	AccountId []int `json:"account_id,omitempty"`
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
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
}

type RatingsDates struct {
	// Dates with available ratings
	Dates []wgnTime.UnixTime `json:"dates,omitempty"`
}
