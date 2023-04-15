// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClanmessagesLikesOptions struct {
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
}

type ClanmessagesLikes struct {
	// Liker Account ID
	AccountId *int `json:"account_id,omitempty"`
	// Date when message was liked
	LikedAt *wgnTime.UnixTime `json:"liked_at,omitempty"`
}
