package wows

type EncyclopediaBattlearenas struct {
	// Map name
	Name string `json:"name,omitempty"`
	// URL to the map icon
	Icon string `json:"icon,omitempty"`
	// Map description
	Description string `json:"description,omitempty"`
	// Map ID
	BattleArenaId int `json:"battle_arena_id,omitempty"`
} 
