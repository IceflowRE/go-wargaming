package wotb

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// ClanmessagesLikesOptions options.
type ClanmessagesLikesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	//
	// "en" - English
	// "ru" - Русский (by default)
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
	Language *string
}

type ClanmessagesLikes struct {
	// Liker Account ID
	AccountId *int `json:"account_id,omitempty"`
	// Date when message was liked
	LikedAt *wgnTime.UnixTime `json:"liked_at,omitempty"`
}
