package wgn

import (
	"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"
)

// WargagContentOptions options.
type WargagContentOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string
	// Publication author ID
	AccountId *int
	// Content category ID
	CategoryId *int
	// Content ID. When this parameter is specified, all other passed-in parameters are ignored.
	ContentId *int
	// Publication date
	Date *wgnTime.UnixTime
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Sorting. Default is "-date". Valid values:
	// 
	// "date" - by publication date 
	// "-date" - by publication date in reverse order (by default)
	// "rating" - by rating value 
	// "-rating" - by rating value in reverse order
	OrderBy *string
	// Result page number
	PageNo *int
	// Threshold of publication rating
	RatingThreshold *int
	// Tag ID
	TagId *int
	// Content type. Valid values:
	// 
	// "quote" - Quote content 
	// "video" - Video content 
	// "picture" - Image content type
	Type_ *string
}

type WargagContent struct {
	// Publication author ID
	AccountId *int `json:"account_id,omitempty"`
	// Indicates the possibility to vote for content. This data requires a valid access_token for the specified account.
	AllowedToVote *bool `json:"allowed_to_vote,omitempty"`
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
