// Auto generated file!

package wotb

type ClansGlossaryOptions struct {
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

type ClansGlossary struct {
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
	// Clan settings
	Settings *struct {
		// Max number of clan members
		MaxMembersCount *int `json:"max_members_count,omitempty"`
	} `json:"settings,omitempty"`
}
