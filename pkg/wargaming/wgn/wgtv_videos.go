package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type WgtvVideos struct {
	// Youtube link
	VideoUrl string `json:"video_url,omitempty"`
	// Youtube ID
	VideoId string `json:"video_id,omitempty"`
	// List of IDs of vehicles covered in the video
	Vehicles map[string]string `json:"vehicles,omitempty"`
	// Localized name of the video
	Title string `json:"title,omitempty"`
	// List of images
	Thumbnails struct {
		// Width
		Width int `json:"width,omitempty"`
		// Image link
		Url string `json:"url,omitempty"`
		// Height
		Height int `json:"height,omitempty"`
	} `json:"thumbnails,omitempty"`
	// Publication date and time
	PublishedAt wgnTime.UnixTime `json:"published_at,omitempty"`
	// List of project IDs
	ProjectId []int `json:"project_id,omitempty"`
	// Program ID
	ProgramId int `json:"program_id,omitempty"`
	// "Important" mark
	Important bool `json:"important,omitempty"`
	// Name of the video on YouTube
	ExtTitle string `json:"ext_title,omitempty"`
	// Video duration in seconds. The field can return null.
	Duration int `json:"duration,omitempty"`
	// Localized description of the video
	Description string `json:"description,omitempty"`
	// List of category IDs
	CategoryId []int `json:"category_id,omitempty"`
} 
