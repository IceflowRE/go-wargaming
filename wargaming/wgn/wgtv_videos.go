// Auto generated file!

package wgn

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// WgtvVideosOptions options.
type WgtvVideosOptions struct {
	// Content category ID. Maximum limit: 100.
	CategoryId []int
	// Dated from the specified date
	DateFrom *wgnTime.UnixTime
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// "Important" mark. Min value is 0. Maximum value: 1.
	Important *int
	// Localization language. Default is "ru". Valid values:
	//
	// "en" - English
	// "ru" - Русский (by default)
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string
	// Number of returned entries (fewer can be returned, but not more than 1000). If the limit sent exceeds 1000, a limit of 100 is applied (by default).
	Limit *int
	// Result page number. Default is 1. Min value is 1.
	PageNo *int
	// Program ID. Maximum limit: 100.
	ProgramId []int
	// Game project ID. Maximum limit: 100.
	ProjectId []int
	// Text for search by name
	Q *string
	// Vehicle ID. Maximum limit: 100.
	VehicleId []int
	// Youtube ID. Maximum limit: 100.
	VideoId []string
}

type WgtvVideos struct {
	// List of category IDs
	CategoryId []int `json:"category_id,omitempty"`
	// Localized description of the video
	Description *string `json:"description,omitempty"`
	// Video duration in seconds. The field can return null.
	Duration *int `json:"duration,omitempty"`
	// Name of the video on YouTube
	ExtTitle *string `json:"ext_title,omitempty"`
	// "Important" mark
	Important *bool `json:"important,omitempty"`
	// Program ID
	ProgramId *int `json:"program_id,omitempty"`
	// List of project IDs
	ProjectId []int `json:"project_id,omitempty"`
	// Publication date and time
	PublishedAt *wgnTime.UnixTime `json:"published_at,omitempty"`
	// List of images
	Thumbnails *struct {
		// Height
		Height *int `json:"height,omitempty"`
		// Image link
		Url *string `json:"url,omitempty"`
		// Width
		Width *int `json:"width,omitempty"`
	} `json:"thumbnails,omitempty"`
	// Localized name of the video
	Title *string `json:"title,omitempty"`
	// List of IDs of vehicles covered in the video
	Vehicles map[string]string `json:"vehicles,omitempty"`
	// Youtube ID
	VideoId *string `json:"video_id,omitempty"`
	// Youtube link
	VideoUrl *string `json:"video_url,omitempty"`
}
