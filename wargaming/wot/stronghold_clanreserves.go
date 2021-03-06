package wot

// StrongholdClanreservesOptions options.
type StrongholdClanreservesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	//
	// "ru" - Russian (by default)
	// "be" - Belarusian
	// "uk" - Ukrainian
	// "kk" - Kazakh
	Language *string
}

type StrongholdClanreserves struct {
	// Reserve efficiency type
	BonusType *string `json:"bonus_type,omitempty"`
	// Indicates if the Reserve is a One-Time-Effect Reserve
	Disposable *bool `json:"disposable,omitempty"`
	// URL to icon
	Icon *string `json:"icon,omitempty"`
	// Available clan Reserves and their status
	InStock *struct {
		// Duration of clan Reserves of each level
		ActionTime *int `json:"action_time,omitempty"`
		// Activation time of clan Reserves of each level
		ActivatedAt *int `json:"activated_at,omitempty"`
		// Expiration time of clan Reserves of each level
		ActiveTill *int `json:"active_till,omitempty"`
		// Number of clan Reserves of each level
		Amount *int `json:"amount,omitempty"`
		// Reserve efficiencies
		BonusValues *struct {
			// Battle type
			BattleType *string `json:"battle_type,omitempty"`
			// Reserve efficiency for each battle type
			Value *float32 `json:"value,omitempty"`
		} `json:"bonus_values,omitempty"`
		// Level of available clan Reserves
		Level *int `json:"level,omitempty"`
		// Status of clan Reserves of each level
		Status *string `json:"status,omitempty"`
		// Indicates if the Reserve is only for Tier X vehicles
		XLevelOnly *bool `json:"x_level_only,omitempty"`
	} `json:"in_stock,omitempty"`
	// Reserve name
	Name *string `json:"name,omitempty"`
	// Reserve type
	Type_ *string `json:"type,omitempty"`
}
