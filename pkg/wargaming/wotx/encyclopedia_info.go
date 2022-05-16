package wotx

type EncyclopediaInfo struct {
	// Available vehicle types
	VehicleTypes map[string]string `json:"vehicle_types,omitempty"`
	// Nations available
	VehicleNations map[string]string `json:"vehicle_nations,omitempty"`
	// Game client version
	GameVersion string `json:"game_version,omitempty"`
	// Award sections
	AchievementSections struct {
		// Award section order
		Order int `json:"order,omitempty"`
		// Award section name
		Name string `json:"name,omitempty"`
	} `json:"achievement_sections,omitempty"`
} 
