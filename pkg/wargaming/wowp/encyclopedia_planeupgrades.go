package wowp

type EncyclopediaPlaneupgrades struct {
	// Localized slot_name field
	SlotNameI18N string `json:"slot_name_i18n,omitempty"`
	// Slot name
	SlotName string `json:"slot_name,omitempty"`
	// Modules available
	Modules struct {
		// Parent module ID in Tech Tree
		ParentId int `json:"parent_id,omitempty"`
		// Name of parent module in Tech Tree
		Parent string `json:"parent,omitempty"`
		// Aircraft ID to be opened with module research
		NextPlaneId int `json:"next_plane_id,omitempty"`
		// Module name
		Name string `json:"name,omitempty"`
		// Module ID
		ModuleId int `json:"module_id,omitempty"`
		// Indicates if the module is standard
		IsDefault bool `json:"is_default,omitempty"`
		// List of IDs of incompatible modules
		IncompatibleModules []int `json:"incompatible_modules,omitempty"`
		// Number of installed modules
		Count int `json:"count,omitempty"`
		// Incompatible modules
		Conflicts []string `json:"conflicts,omitempty"`
		// ID of unique bind of slot and module
		BindId int `json:"bind_id,omitempty"`
	} `json:"modules,omitempty"`
} 
