package wot

type AccountAchievements struct {
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Achievement progress
	Frags map[string]string `json:"frags,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
} 
