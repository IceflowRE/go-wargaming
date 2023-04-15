// Auto generated file!

package wows

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type EncyclopediaInfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
}

type EncyclopediaInfo struct {
	// Game client version
	GameVersion *string `json:"game_version,omitempty"`
	// List of languages supported by encyclopedia methods
	Languages map[string]string `json:"languages,omitempty"`
	// Types of Modifications available
	ShipModifications map[string]string `json:"ship_modifications,omitempty"`
	// Types of modules available
	ShipModules map[string]string `json:"ship_modules,omitempty"`
	// Nations available
	ShipNations map[string]string `json:"ship_nations,omitempty"`
	// Images of ship types
	ShipTypeImages *struct {
		// Ship type image
		Image *string `json:"image,omitempty"`
		// Elite ships icon
		ImageElite *string `json:"image_elite,omitempty"`
		// Premium ships icon
		ImagePremium *string `json:"image_premium,omitempty"`
	} `json:"ship_type_images,omitempty"`
	// Types of ships available
	ShipTypes map[string]string `json:"ship_types,omitempty"`
	// Time when details on ships were updated
	ShipsUpdatedAt *wgnTime.UnixTime `json:"ships_updated_at,omitempty"`
}
