package wows

type AccountAchievements struct {
	// Achievement progress
	Progress map[string]string `json:"progress,omitempty"`
	// Battle achievements earned
	Battle map[string]string `json:"battle,omitempty"`
} 
