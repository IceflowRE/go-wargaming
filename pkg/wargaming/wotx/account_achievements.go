package wotx

type AccountAchievements struct {
	// Ribbons earned
	Ribbons map[string]string `json:"ribbons,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
} 
