package wot

type TanksAchievements struct {
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Current values of Achievement Series
	Series map[string]string `json:"series,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
