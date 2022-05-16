package wot

type RatingsTypes struct {
	// Rating period
	Type_ string `json:"type,omitempty"`
	// Rating threshold
	Threshold int `json:"threshold,omitempty"`
	// Rating categories
	RankFields []string `json:"rank_fields,omitempty"`
} 
