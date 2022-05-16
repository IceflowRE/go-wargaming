package wotb

type AccountAchievements struct {
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
} 
