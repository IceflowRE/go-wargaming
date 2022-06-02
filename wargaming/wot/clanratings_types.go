package wot


type ClanratingsTypes struct {
	// Rating categories
	RankFields []string `json:"rank_fields,omitempty"`
	// Rating period
	Type_ *string `json:"type,omitempty"`
}
