package wot

type EncyclopediaCrewroles struct {
	// List of crew member qualifications
	Skills []string `json:"skills,omitempty"`
	// Сrew qualification ID
	Role string `json:"role,omitempty"`
	// Сrew qualification name
	Name string `json:"name,omitempty"`
} 
