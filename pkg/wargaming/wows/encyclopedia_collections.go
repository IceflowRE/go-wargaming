package wows

type EncyclopediaCollections struct {
	// Collection tag
	Tag string `json:"tag,omitempty"`
	// Collection name
	Name string `json:"name,omitempty"`
	// Link to collection image
	Image string `json:"image,omitempty"`
	// Collection description
	Description string `json:"description,omitempty"`
	// Collection ID
	CollectionId int `json:"collection_id,omitempty"`
	// Collection item cost in duplicates
	CardCost int `json:"card_cost,omitempty"`
} 
