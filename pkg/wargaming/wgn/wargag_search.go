package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type WargagSearch struct {
	// Link to original publication
	WargagUrl string `json:"wargag_url,omitempty"`
	// Content type
	Type_ string `json:"type,omitempty"`
	// Tag ID
	TagId int `json:"tag_id,omitempty"`
	// Publication topic
	Subject string `json:"subject,omitempty"`
	// Current rating
	Rating int `json:"rating,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Video/image link
	MediaUrl string `json:"media_url,omitempty"`
	// Link to image preview. Available for picture content only.
	MediaPreviewUrl string `json:"media_preview_url,omitempty"`
	// Publication text
	Description string `json:"description,omitempty"`
	// Publication date in UNIX timestamp or ISO 8601
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Publication ID
	ContentId int `json:"content_id,omitempty"`
	// Content category ID
	CategoryId int `json:"category_id,omitempty"`
	// Content author
	Author struct {
		// Title icon
		StatusImage string `json:"status_image,omitempty"`
		// Author's title
		Status string `json:"status,omitempty"`
		// Author's reputation
		Reputation int `json:"reputation,omitempty"`
	} `json:"author,omitempty"`
	// Indicates the possibility to vote for content. This data requires a valid access_token for the specified account.
	AllowedToVote bool `json:"allowed_to_vote,omitempty"`
	// Publication author ID
	AccountId int `json:"account_id,omitempty"`
} 
