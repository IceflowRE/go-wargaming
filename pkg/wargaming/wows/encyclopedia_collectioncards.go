package wows

type EncyclopediaCollectioncards struct {
	// Item tag
	Tag string `json:"tag,omitempty"`
	// Item name
	Name string `json:"name,omitempty"`
	// Item card icons
	Images struct {
		// URL to the small card icon
		Small string `json:"small,omitempty"`
		// URL to the medium card icon
		Medium string `json:"medium,omitempty"`
		// URL to the large card icon
		Large string `json:"large,omitempty"`
	} `json:"images,omitempty"`
	// Item description
	Description string `json:"description,omitempty"`
	// Collection ID
	CollectionId int `json:"collection_id,omitempty"`
	// Item ID
	CardId int `json:"card_id,omitempty"`
} 
