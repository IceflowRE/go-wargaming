// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClanmessagesMessagesOptions struct {
	// Search for messages with end date of relevance equal or after the value. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
	ExpiresAfter *wgnTime.UnixTime `json:"expires_after,omitempty"`
	// Search for messages with end date of relevance before the value. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
	ExpiresBefore *wgnTime.UnixTime `json:"expires_before,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Message importance. Valid values:
	//
	// "important" - Important messages
	// "standard" - Standard messages
	Importance *string `json:"importance,omitempty"`
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
	// Number of returned entries. If the value sent exceeds 100, a limit of 25 is applied (by default). Default is 25. Min value is 1.
	Limit *int `json:"limit,omitempty"`
	// Message ID
	MessageId *int `json:"message_id,omitempty"`
	// Sorting. Valid values:
	//
	//
	// importance - by message importance
	//
	//
	// -importance - by message importance in reverse order
	//
	//
	// created_at - by date of message creation
	//
	//
	// -created_at - by date of message creation in reverse order
	//
	//
	// updated_at - by message update date
	//
	//
	// -updated_at - by message update date in reverse order
	//
	//
	// type - by message type
	//
	//
	// -type - by message type in reverse order. Default is "importance, type". Maximum limit: 100.
	//
	OrderBy []string `json:"order_by,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
	// Message status. Valid values:
	//
	// "active" - Active message
	// "deleted" - Deleted message
	Status *string `json:"status,omitempty"`
	// Message type. Valid values:
	//
	// "general" - General messages
	// "training" - Training messages
	// "meeting" - Meeting messages
	// "battle" - Battle messages
	Type *string `json:"type,omitempty"`
}

type ClanmessagesMessages struct {
	// Message author ID
	AuthorId *int `json:"author_id,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Message creation date
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// ID of a player who has changed the message
	EditorId *int `json:"editor_id,omitempty"`
	// Date when message will become irrelevant
	ExpiresAt *wgnTime.UnixTime `json:"expires_at,omitempty"`
	// Message importance
	Importance *string `json:"importance,omitempty"`
	// Number of likes
	Likes *int `json:"likes,omitempty"`
	// Message text
	Message *string `json:"message,omitempty"`
	// Message ID
	MessageId *int `json:"message_id,omitempty"`
	// Message title
	Title *string `json:"title,omitempty"`
	// Message type
	Type *string `json:"type,omitempty"`
	// Date when message was updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
