// Auto generated file!

package wgn

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

type WargagNewcomment struct {
	// Comment author ID
	AccountId *int `json:"account_id,omitempty"`
	// Comment Author
	Author *struct {
		// Author's reputation
		Reputation *int `json:"reputation,omitempty"`
		// Author's title
		Status *string `json:"status,omitempty"`
		// Title icon
		StatusImage *string `json:"status_image,omitempty"`
	} `json:"author,omitempty"`
	// Comment text
	Comment *string `json:"comment,omitempty"`
	// Comment ID
	CommentId *int `json:"comment_id,omitempty"`
	// Publication ID
	ContentId *int `json:"content_id,omitempty"`
	// Comment date
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
}
