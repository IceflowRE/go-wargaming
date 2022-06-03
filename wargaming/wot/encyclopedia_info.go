package wot

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// EncyclopediaInfoOptions options.
type EncyclopediaInfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
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
}

type EncyclopediaInfo struct {
	// Award sections
	AchievementSections *struct {
		// Award section name
		Name *string `json:"name,omitempty"`
		// Award section order
		Order *int `json:"order,omitempty"`
	} `json:"achievement_sections,omitempty"`
	// Game client version
	GameVersion *string `json:"game_version,omitempty"`
	// List of supported languages
	Languages map[string]string `json:"languages,omitempty"`
	// Time when vehicles details in Tankopedia were updated
	TanksUpdatedAt *wgnTime.UnixTime `json:"tanks_updated_at,omitempty"`
	// Available crew qualifications
	VehicleCrewRoles map[string]string `json:"vehicle_crew_roles,omitempty"`
	// Nations available
	VehicleNations map[string]string `json:"vehicle_nations,omitempty"`
	// Available vehicle types
	VehicleTypes map[string]string `json:"vehicle_types,omitempty"`
}
