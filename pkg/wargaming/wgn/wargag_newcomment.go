package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type WargagNewcomment struct {
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Comment date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Publication ID
	ContentId int `json:"content_id,omitempty"`
	// Comment ID
	CommentId int `json:"comment_id,omitempty"`
	// Comment text
	Comment string `json:"comment,omitempty"`
	// Comment Author
	Author struct {
		// Title icon
		StatusImage string `json:"status_image,omitempty"`
		// Author's title
		Status string `json:"status,omitempty"`
		// Author's reputation
		Reputation int `json:"reputation,omitempty"`
	} `json:"author,omitempty"`
	// Comment author ID
	AccountId int `json:"account_id,omitempty"`
} 
