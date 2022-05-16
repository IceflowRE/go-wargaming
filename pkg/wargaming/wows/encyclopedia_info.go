package wows

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type EncyclopediaInfo struct {
	// Time when details on ships were updated
	ShipsUpdatedAt wgnTime.UnixTime `json:"ships_updated_at,omitempty"`
	// Types of ships available
	ShipTypes map[string]string `json:"ship_types,omitempty"`
	// Images of ship types
	ShipTypeImages struct {
		// Premium ships icon
		ImagePremium string `json:"image_premium,omitempty"`
		// Elite ships icon
		ImageElite string `json:"image_elite,omitempty"`
		// Ship type image
		Image string `json:"image,omitempty"`
	} `json:"ship_type_images,omitempty"`
	// Nations available
	ShipNations map[string]string `json:"ship_nations,omitempty"`
	// Types of modules available
	ShipModules map[string]string `json:"ship_modules,omitempty"`
	// Types of Modifications available
	ShipModifications map[string]string `json:"ship_modifications,omitempty"`
	// List of languages supported by encyclopedia methods
	Languages map[string]string `json:"languages,omitempty"`
	// Game client version
	GameVersion string `json:"game_version,omitempty"`
} 
