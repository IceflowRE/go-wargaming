package wgn

type ServersInfo struct {
	// Server ID
	Server string `json:"server,omitempty"`
	// Number of online players
	PlayersOnline int `json:"players_online,omitempty"`
} 
