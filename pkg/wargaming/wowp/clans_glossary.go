package wowp

type ClansGlossary struct {
	// Clan settings
	Settings struct {
		// Maximum clan members
		MaxMembersCount int `json:"max_members_count,omitempty"`
	} `json:"settings,omitempty"`
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
} 
