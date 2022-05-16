package wows

type EncyclopediaCrewranks struct {
	// Rank
	Rank int `json:"rank,omitempty"`
	// Rank Name by subnation index
	Names map[string]string `json:"names,omitempty"`
	// Rank name
	Name string `json:"name,omitempty"`
	// Experience
	Experience int `json:"experience,omitempty"`
} 
