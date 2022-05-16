package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansList struct {
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Game where clan was created
	Game string `json:"game,omitempty"`
	// Information on clan emblems in games and on clan portal
	Emblems struct {
		// List of links to 64x64 px emblems
		X64 map[string]string `json:"x64,omitempty"`
		// List of links to 32x32 px emblems
		X32 map[string]string `json:"x32,omitempty"`
		// List of links to 256x256 px emblems
		X256 map[string]string `json:"x256,omitempty"`
		// List of links to 24x24 px emblems
		X24 map[string]string `json:"x24,omitempty"`
		// List of links to 195x195 px emblems
		X195 map[string]string `json:"x195,omitempty"`
	} `json:"emblems,omitempty"`
	// Clan creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
} 
