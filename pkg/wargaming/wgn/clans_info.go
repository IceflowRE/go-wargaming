package wgn

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClansInfo struct {
	// Time when clan details were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt wgnTime.UnixTime `json:"renamed_at,omitempty"`
	// Restricted clan information
	Private struct {
		// Amount of gold in the сlan Treasury
		Treasury int `json:"treasury,omitempty"`
		// List of clan members with an active game session in World of Tanks.
		// An extra field.
		OnlineMembers []int `json:"online_members,omitempty"`
		// Clan Treasury
		ClanTreasury struct {
			// Amount of gold in the сlan Treasury
			Gold int `json:"gold,omitempty"`
			// Number of bonds in the сlan Treasury
			Crystal int `json:"crystal,omitempty"`
			// Amount of credits in the сlan Treasury
			Credits int `json:"credits,omitempty"`
		} `json:"clan_treasury,omitempty"`
	} `json:"private,omitempty"`
	// Old clan tag
	OldTag string `json:"old_tag,omitempty"`
	// Old clan name
	OldName string `json:"old_name,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan motto
	Motto string `json:"motto,omitempty"`
	// Number of clan members
	MembersCount int `json:"members_count,omitempty"`
	// Information on clan members. Field format depends on members_key input parameter.
	Members struct {
		// Position
		RoleI18N string `json:"role_i18n,omitempty"`
		// Technical position name
		Role string `json:"role,omitempty"`
		// Date when player joined clan
		JoinedAt wgnTime.UnixTime `json:"joined_at,omitempty"`
		// Player name
		AccountName string `json:"account_name,omitempty"`
		// Player account ID
		AccountId int `json:"account_id,omitempty"`
	} `json:"members,omitempty"`
	// Commander's name
	LeaderName string `json:"leader_name,omitempty"`
	// Clan Commander ID
	LeaderId int `json:"leader_id,omitempty"`
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded bool `json:"is_clan_disbanded,omitempty"`
	// Game where clan was created
	Game string `json:"game,omitempty"`
	// Information on clan emblems in games and on clan portal
	Emblems struct {
		// List of links to 64x64 px emblems
		X64 map[string]string `json:"x64,omitempty"`
		// List of links to 32x32 px emblems
		X32 map[string]string `json:"x32,omitempty"`
		// List of links to 256x256 px emblems
		X256 map[string]string `json:"x256,omitempty"`
		// List of links to 24x24 px emblems
		X24 map[string]string `json:"x24,omitempty"`
		// List of links to 195x195 px emblems
		X195 map[string]string `json:"x195,omitempty"`
	} `json:"emblems,omitempty"`
	// Clan description in HTML
	DescriptionHtml string `json:"description_html,omitempty"`
	// Clan description
	Description string `json:"description,omitempty"`
	// Clan creator's name
	CreatorName string `json:"creator_name,omitempty"`
	// Clan creator ID
	CreatorId int `json:"creator_id,omitempty"`
	// Clan creation date
	CreatedAt wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan color in HEX #RRGGBB
	Color string `json:"color,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Clan can invite players
	AcceptsJoinRequests bool `json:"accepts_join_requests,omitempty"`
} 
