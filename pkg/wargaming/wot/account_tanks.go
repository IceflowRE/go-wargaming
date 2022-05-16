package wot

type AccountTanks struct {
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Vehicle statistics
	Statistics struct {
		// Victories
		Wins int `json:"wins,omitempty"`
		// Battles fought
		Battles int `json:"battles,omitempty"`
	} `json:"statistics,omitempty"`
	// Mastery Badges:
	// 
	// 0 — None
	// 1 — 3rd Class 
	// 2 — 2nd Class
	// 3 — 1st Class
	// 4 — Ace Tanker
	MarkOfMastery int `json:"mark_of_mastery,omitempty"`
} 
