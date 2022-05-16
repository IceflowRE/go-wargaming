package wotb

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClanmessagesMessages struct {
	// Date when message was updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Message type
	Type_ string `json:"type,omitempty"`
	// Message title
	Title string `json:"title,omitempty"`
	// Message ID
	MessageId int `json:"message_id,omitempty"`
	// Message text
	Message string `json:"message,omitempty"`
	// Number of likes
	Likes int `json:"likes,omitempty"`
	// Message importance
	Importance string `json:"importance,omitempty"`
	// Date when message will become irrelevant
	ExpiresAt wgnTime.UnixTime `json:"expires_at,omitempty"`
	// ID of a player who has changed the message
	EditorId int `json:"editor_id,omitempty"`
	// Message creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Message author ID
	AuthorId int `json:"author_id,omitempty"`
} 
