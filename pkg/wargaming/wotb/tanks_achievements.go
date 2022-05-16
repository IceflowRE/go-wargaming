package wotb

type TanksAchievements struct {
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Maximum values of achievement series
	MaxSeries map[string]string `json:"max_series,omitempty"`
	// Achievements earned
	Achievements map[string]string `json:"achievements,omitempty"`
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
} 
