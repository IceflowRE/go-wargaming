package wows

type EncyclopediaAccountlevels struct {
	// Service Record level
	Tier int `json:"tier,omitempty"`
	// Points to achieve the Service Record level
	Points int `json:"points,omitempty"`
	// URL to Service Record level icon
	Image string `json:"image,omitempty"`
} 
