package wows

type AccountList struct {
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Player ID
	AccountId int `json:"account_id,omitempty"`
} 
