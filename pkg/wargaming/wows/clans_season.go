package wows

type ClansSeason struct {
	// Season start time
	StartTime int `json:"start_time,omitempty"`
	// Minimal ship Tier in a season
	ShipTierMin int `json:"ship_tier_min,omitempty"`
	// Maximum ship Tier in a season
	ShipTierMax int `json:"ship_tier_max,omitempty"`
	// Season ID
	SeasonId int `json:"season_id,omitempty"`
	// Season name
	Name string `json:"name,omitempty"`
	// Leagues in the season Rating
	Leagues struct {
		// League name
		Name string `json:"name,omitempty"`
		// League icon
		Icon string `json:"icon,omitempty"`
		// Color of the league icon
		Color string `json:"color,omitempty"`
	} `json:"leagues,omitempty"`
	// Season end time
	FinishTime int `json:"finish_time,omitempty"`
	// Number of the Rating points required for progression to the next league
	DivisionPoints int `json:"division_points,omitempty"`
} 
