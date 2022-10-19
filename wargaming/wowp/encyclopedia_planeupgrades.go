// Auto generated file!

package wowp

type EncyclopediaPlaneupgradesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
}

type EncyclopediaPlaneupgrades struct {
	// Modules available
	Modules *struct {
		// ID of unique bind of slot and module
		BindId *int `json:"bind_id,omitempty"`
		// Incompatible modules
		Conflicts []string `json:"conflicts,omitempty"`
		// Number of installed modules
		Count *int `json:"count,omitempty"`
		// List of IDs of incompatible modules
		IncompatibleModules []int `json:"incompatible_modules,omitempty"`
		// Indicates if the module is standard
		IsDefault *bool `json:"is_default,omitempty"`
		// Module ID
		ModuleId *int `json:"module_id,omitempty"`
		// Module name
		Name *string `json:"name,omitempty"`
		// Aircraft ID to be opened with module research
		NextPlaneId *int `json:"next_plane_id,omitempty"`
		// Name of parent module in Tech Tree
		Parent *string `json:"parent,omitempty"`
		// Parent module ID in Tech Tree
		ParentId *int `json:"parent_id,omitempty"`
	} `json:"modules,omitempty"`
	// Slot name
	SlotName *string `json:"slot_name,omitempty"`
	// Localized slot_name field
	SlotNameI18n *string `json:"slot_name_i18n,omitempty"`
}
