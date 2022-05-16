package wot

type EncyclopediaPersonalmissions struct {
	// Campaign operations
	Operations struct {
		// Number of branches until the next operation
		SetsToNext int `json:"sets_to_next,omitempty"`
		// Number of mission branches of the operation
		SetsCount int `json:"sets_count,omitempty"`
		// Reward for operation
		Reward struct {
			// Vehicle IDs
			Tanks []int `json:"tanks,omitempty"`
			// Slots
			Slots int `json:"slots,omitempty"`
		} `json:"reward,omitempty"`
		// Operation ID
		OperationId int `json:"operation_id,omitempty"`
		// Next operation ID
		NextId int `json:"next_id,omitempty"`
		// Operation name
		Name string `json:"name,omitempty"`
		// Number of missions in the branch
		MissionsInSet int `json:"missions_in_set,omitempty"`
		// Operation missions
		Missions struct {
			// Mission tags
			Tags []string `json:"tags,omitempty"`
			// Mission branch ID
			SetId int `json:"set_id,omitempty"`
			// Rewards grouped by mission conditions
			Rewards struct {
				// Tokens
				Tokens int `json:"tokens,omitempty"`
				// Slots
				Slots int `json:"slots,omitempty"`
				// Days of Premium Account
				Premium int `json:"premium,omitempty"`
				// List of equipment or consumables in the following format: Id–number of items
				Items map[string]string `json:"items,omitempty"`
				// Free Experience
				FreeXp int `json:"free_xp,omitempty"`
				// Credits
				Credits int `json:"credits,omitempty"`
				// Mission conditions
				Conditions string `json:"conditions,omitempty"`
				// Bunks in Barracks
				Berths int `json:"berths,omitempty"`
			} `json:"rewards,omitempty"`
			// Mission name
			Name string `json:"name,omitempty"`
			// Mission ID
			MissionId int `json:"mission_id,omitempty"`
			// Minimum vehicle Tier
			MinLevel int `json:"min_level,omitempty"`
			// Maximum vehicle Tier
			MaxLevel int `json:"max_level,omitempty"`
			// Hints
			Hint string `json:"hint,omitempty"`
			// Mission description
			Description string `json:"description,omitempty"`
		} `json:"missions,omitempty"`
		// Link to an operation image
		Image string `json:"image,omitempty"`
		// Operation description
		Description string `json:"description,omitempty"`
	} `json:"operations,omitempty"`
	// Campaign name
	Name string `json:"name,omitempty"`
	// Campaign description
	Description string `json:"description,omitempty"`
	// Campaign ID
	CampaignId int `json:"campaign_id,omitempty"`
} 
