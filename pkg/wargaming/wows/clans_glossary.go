package wows

type ClansGlossary struct {
	// Clan settings
	Settings struct {
		// Max number of clan members
		MaxMembersCount int `json:"max_members_count,omitempty"`
	} `json:"settings,omitempty"`
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
	// Installations
	Buildings struct {
		// Type of ships that will get a bonus after building this installation. If value = "empty string" or "null", the bonus will be received on ships of all types.
		ShipType string `json:"ship_type,omitempty"`
		// Tier of ships that will get a bonus after building this installation. If value = "null", the bonus will be received on ships of all Tiers.
		ShipTier int `json:"ship_tier,omitempty"`
		// The nation of ships that will get a bonus from this installation. If value = "empty string" or "null", the bonus will be received on ships of all nations.
		ShipNation string `json:"ship_nation,omitempty"`
		// Installation name
		Name string `json:"name,omitempty"`
		// Building cost in oil
		Cost int `json:"cost,omitempty"`
		// Installation ID
		BuildingTypeId int `json:"building_type_id,omitempty"`
		// Installation ID
		BuildingId int `json:"building_id,omitempty"`
		// The value of the bonus based on this bonus type.
		BonusValue int `json:"bonus_value,omitempty"`
		// The type of the bonus that is provided to the clan members after building the installation. Existing bonus types:
		// 
		// exp_boost—rate of additional XP;
		// members_count—number of clan members on which the clan size is increased;
		// maintenance_discount—rate of reducing cost of servicing ships;
		// purchase_discount—rate of reducing cost of researched ships;
		// dummy—no bonuses.
		BonusType string `json:"bonus_type,omitempty"`
	} `json:"buildings,omitempty"`
	// Installation type
	BuildingTypes struct {
		// Structure name
		Name string `json:"name,omitempty"`
		// Installation ID
		BuildingTypeId int `json:"building_type_id,omitempty"`
	} `json:"building_types,omitempty"`
} 
