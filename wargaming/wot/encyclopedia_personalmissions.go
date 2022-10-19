// Auto generated file!

package wot

type EncyclopediaPersonalmissionsOptions struct {
	// Campaign ID. Maximum limit: 100.
	CampaignId []int `json:"campaign_id,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
	// Operation ID. Maximum limit: 100.
	OperationId []int `json:"operation_id,omitempty"`
	// Mission branch ID. Maximum limit: 100.
	SetId []int `json:"set_id,omitempty"`
	// Mission tag. Maximum limit: 100.
	Tag []string `json:"tag,omitempty"`
}

type EncyclopediaPersonalmissions struct {
	// Campaign ID
	CampaignId *int `json:"campaign_id,omitempty"`
	// Campaign description
	Description *string `json:"description,omitempty"`
	// Campaign name
	Name *string `json:"name,omitempty"`
	// Campaign operations
	Operations *struct {
		// Operation description
		Description *string `json:"description,omitempty"`
		// Link to an operation image
		Image *string `json:"image,omitempty"`
		// Operation missions
		Missions *struct {
			// Mission description
			Description *string `json:"description,omitempty"`
			// Hints
			Hint *string `json:"hint,omitempty"`
			// Maximum vehicle Tier
			MaxLevel *int `json:"max_level,omitempty"`
			// Minimum vehicle Tier
			MinLevel *int `json:"min_level,omitempty"`
			// Mission ID
			MissionId *int `json:"mission_id,omitempty"`
			// Mission name
			Name *string `json:"name,omitempty"`
			// Rewards grouped by mission conditions
			Rewards *struct {
				// Bunks in Barracks
				Berths *int `json:"berths,omitempty"`
				// Mission conditions
				Conditions *string `json:"conditions,omitempty"`
				// Credits
				Credits *int `json:"credits,omitempty"`
				// Free Experience
				FreeXp *int `json:"free_xp,omitempty"`
				// List of equipment or consumables in the following format: Id–number of items
				Items map[string]string `json:"items,omitempty"`
				// Days of Premium Account
				Premium *int `json:"premium,omitempty"`
				// Slots
				Slots *int `json:"slots,omitempty"`
				// Tokens
				Tokens *int `json:"tokens,omitempty"`
			} `json:"rewards,omitempty"`
			// Mission branch ID
			SetId *int `json:"set_id,omitempty"`
			// Mission tags
			Tags []string `json:"tags,omitempty"`
		} `json:"missions,omitempty"`
		// Number of missions in the branch
		MissionsInSet *int `json:"missions_in_set,omitempty"`
		// Operation name
		Name *string `json:"name,omitempty"`
		// Next operation ID
		NextId *int `json:"next_id,omitempty"`
		// Operation ID
		OperationId *int `json:"operation_id,omitempty"`
		// Reward for operation
		Reward *struct {
			// Slots
			Slots *int `json:"slots,omitempty"`
			// Vehicle IDs
			Tanks []int `json:"tanks,omitempty"`
		} `json:"reward,omitempty"`
		// Number of mission branches of the operation
		SetsCount *int `json:"sets_count,omitempty"`
		// Number of branches until the next operation
		SetsToNext *int `json:"sets_to_next,omitempty"`
	} `json:"operations,omitempty"`
}
