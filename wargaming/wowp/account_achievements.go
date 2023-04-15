// Auto generated file!

package wowp

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AccountAchievementsOptions struct {
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

type AccountAchievements struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Achievements earned
	Achievements *struct {
		// Number
		Count *int `json:"count,omitempty"`
		// First time when an achievement was received
		FirstAt *wgnTime.UnixTime `json:"first_at,omitempty"`
		// Last time when an achievement was received
		LastAt *wgnTime.UnixTime `json:"last_at,omitempty"`
		// Most number
		MaxCount *int `json:"max_count,omitempty"`
	} `json:"achievements,omitempty"`
}
