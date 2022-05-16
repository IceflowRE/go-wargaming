package wot

type StrongholdClanreserves struct {
	// Reserve type
	Type_ string `json:"type,omitempty"`
	// Reserve name
	Name string `json:"name,omitempty"`
	// Available clan Reserves and their status
	InStock struct {
		// Indicates if the Reserve is only for Tier X vehicles
		XLevelOnly bool `json:"x_level_only,omitempty"`
		// Status of clan Reserves of each level
		Status string `json:"status,omitempty"`
		// Level of available clan Reserves
		Level int `json:"level,omitempty"`
		// Reserve efficiencies
		BonusValues struct {
			// Reserve efficiency for each battle type
			Value float32 `json:"value,omitempty"`
			// Battle type
			BattleType string `json:"battle_type,omitempty"`
		} `json:"bonus_values,omitempty"`
		// Number of clan Reserves of each level
		Amount int `json:"amount,omitempty"`
		// Expiration time of clan Reserves of each level
		ActiveTill int `json:"active_till,omitempty"`
		// Activation time of clan Reserves of each level
		ActivatedAt int `json:"activated_at,omitempty"`
		// Duration of clan Reserves of each level
		ActionTime int `json:"action_time,omitempty"`
	} `json:"in_stock,omitempty"`
	// URL to icon
	Icon string `json:"icon,omitempty"`
	// Indicates if the Reserve is a One-Time-Effect Reserve
	Disposable bool `json:"disposable,omitempty"`
	// Reserve efficiency type
	BonusType string `json:"bonus_type,omitempty"`
} 
