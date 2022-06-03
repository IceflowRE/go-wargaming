package wotb

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// ClanmessagesUpdateOptions options.
type ClanmessagesUpdateOptions struct {
	// Date when message will become irrelevant. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00. Date must be in the future.
	ExpiresAt *wgnTime.UnixTime
	// Message importance. Valid values:
	//
	// "important" - Important messages
	// "standard" - Standard messages
	Importance *string
	// Message text. Max. length: 1000. Maximum length: 1000.
	Text *string
	// Message title. Max. length: 100. Maximum length: 100.
	Title *string
	// Message type. Valid values:
	//
	// "general" - General messages
	// "training" - Training messages
	// "meeting" - Meeting messages
	// "battle" - Battle messages
	Type_ *string
}

type ClanmessagesUpdate struct {
	// Message ID
	MessageId *int `json:"message_id,omitempty"`
}
