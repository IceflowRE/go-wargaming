package wows

type EncyclopediaBattletypes struct {
	// Battle type tag
	Tag string `json:"tag,omitempty"`
	// Battle type name
	Name string `json:"name,omitempty"`
	// URL to image
	Image string `json:"image,omitempty"`
	// Battle type description
	Description string `json:"description,omitempty"`
} 
