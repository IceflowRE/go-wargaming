// Auto generated file!

package wotb

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClanmessagesUpdateOptions struct {
	// Date when message will become irrelevant. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00. Date must be in the future.
	ExpiresAt *wgnTime.UnixTime `json:"expires_at,omitempty"`
	// Message importance. Valid values:
	//
	// "important" - Important messages
	// "standard" - Standard messages
	Importance *string `json:"importance,omitempty"`
	// Message text. Max. length: 1000. Maximum length: 1000.
	Text *string `json:"text,omitempty"`
	// Message title. Max. length: 100. Maximum length: 100.
	Title *string `json:"title,omitempty"`
	// Message type. Valid values:
	//
	// "general" - General messages
	// "training" - Training messages
	// "meeting" - Meeting messages
	// "battle" - Battle messages
	Type *string `json:"type,omitempty"`
}

type ClanmessagesUpdate struct {
	// Message ID
	MessageId *int `json:"message_id,omitempty"`
}
