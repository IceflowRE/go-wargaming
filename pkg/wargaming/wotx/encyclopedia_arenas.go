package wotx

type EncyclopediaArenas struct {
	// Localized map name
	NameI18N string `json:"name_i18n,omitempty"`
	// Short map description
	Description string `json:"description,omitempty"`
	// Map type
	CamouflageType string `json:"camouflage_type,omitempty"`
	// Map ID
	ArenaId string `json:"arena_id,omitempty"`
} 
