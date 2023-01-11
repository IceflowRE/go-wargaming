// Auto generated file!

package wot

type StrongholdClanreservesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "pl" - Polish
	// "fr" - French
	// "es" - Spanish
	// "cs" - Czech
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
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
	Type *string `json:"type,omitempty"`
}
