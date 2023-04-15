// Auto generated file!

package wgn

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClansMessageboardOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Name of the game to perform the clan search in. If the parameter is not specified, search will be executed across World of Tanks. Default is "wot". Valid values:
	//
	// "wot" - World of Tanks (by default)
	// "wowp" - World of Warplanes
	Game *string `json:"game,omitempty"`
}

type ClansMessageboard struct {
	// Message author ID
	AuthorId *int `json:"author_id,omitempty"`
	// Message creation date
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// ID of a player who has changed the message
	EditorId *int `json:"editor_id,omitempty"`
	// Indicates if the message has been read
	IsRead *bool `json:"is_read,omitempty"`
	// Message text
	Message *string `json:"message,omitempty"`
	// Date when message was updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
