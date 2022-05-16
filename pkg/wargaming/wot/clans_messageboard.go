package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansMessageboard struct {
	// Date when message was updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Message text
	Message string `json:"message,omitempty"`
	// Indicates if the message has been read
	IsRead bool `json:"is_read,omitempty"`
	// ID of a player who has changed the message
	EditorId int `json:"editor_id,omitempty"`
	// Message creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Message author ID
	AuthorId int `json:"author_id,omitempty"`
} 
