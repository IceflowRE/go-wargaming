// Auto generated file!

package wgn

type WgtvTagsOptions struct {
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
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
}

type WgtvTags struct {
	// List of categories
	Categories []*struct {
		// Content category ID
		CategoryId *int `json:"category_id,omitempty"`
		// Localized name of the category
		Name *string `json:"name,omitempty"`
		// Sort order
		Order *int `json:"order,omitempty"`
	} `json:"categories,omitempty"`
	// List of programs
	Programs []*struct {
		// Localized name of the program
		Name *string `json:"name,omitempty"`
		// Sort order
		Order *int `json:"order,omitempty"`
		// Program ID
		ProgramId *int `json:"program_id,omitempty"`
	} `json:"programs,omitempty"`
	// List of projects
	Projects []*struct {
		// Name of the game project
		Name *string `json:"name,omitempty"`
		// Sort order
		Order *int `json:"order,omitempty"`
		// Game project ID
		ProjectId *int `json:"project_id,omitempty"`
	} `json:"projects,omitempty"`
}
