package wgn

type ClansGlossary struct {
	// Game with clan entities
	Game string `json:"game,omitempty"`
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
} 
