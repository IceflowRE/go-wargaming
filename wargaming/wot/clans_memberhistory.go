// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v3/wargaming/wgnTime"
)

type ClansMemberhistoryOptions struct {
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
type ClansMemberhistory struct {
	// Player account ID
	AccountId *int `json:"account_id,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Date when player joined clan
	JoinedAt *wgnTime.UnixTime `json:"joined_at,omitempty"`
	// Date when player left clan
	LeftAt *wgnTime.UnixTime `json:"left_at,omitempty"`
	// Last position in clan
	Role *string `json:"role,omitempty"`
}
