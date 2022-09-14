// Auto generated file!

package wgn

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

type WargagRate struct {
	// Publication author ID
	AccountId *int `json:"account_id,omitempty"`
	// Content author
	Author *struct {
		// Author's reputation
		Reputation *int `json:"reputation,omitempty"`
		// Author's title
		Status *string `json:"status,omitempty"`
		// Title icon
		StatusImage *string `json:"status_image,omitempty"`
	} `json:"author,omitempty"`
	// Content category ID
	CategoryId *int `json:"category_id,omitempty"`
	// Publication ID
	ContentId *int `json:"content_id,omitempty"`
	// Publication date in UNIX timestamp or ISO 8601
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Publication text
	Description *string `json:"description,omitempty"`
	// Link to image preview. Available for picture content only.
	MediaPreviewUrl *string `json:"media_preview_url,omitempty"`
	// Video/image link
	MediaUrl *string `json:"media_url,omitempty"`
	// Player name
	Nickname *string `json:"nickname,omitempty"`
	// Current rating
	Rating *int `json:"rating,omitempty"`
	// Publication topic
	Subject *string `json:"subject,omitempty"`
	// Tag ID
	TagId *int `json:"tag_id,omitempty"`
	// Content type
	Type_ *string `json:"type,omitempty"`
	// Link to original publication
	WargagUrl *string `json:"wargag_url,omitempty"`
}
