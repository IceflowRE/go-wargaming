package wotx

type ClansGlossary struct {
	// Clan settings
	Settings struct {
		// Clan renaming cost in gold
		RenameFeeGold int `json:"rename_fee_gold,omitempty"`
		// Cooldown for clan renaming (in seconds)
		RenameCooldown int `json:"rename_cooldown,omitempty"`
		// Minimum battles required to create clan
		MinBattlesCount int `json:"min_battles_count,omitempty"`
		// Max number of clan members
		MaxMembersCount int `json:"max_members_count,omitempty"`
		// Maximum number of invitations
		MaxInvitesPerClan int `json:"max_invites_per_clan,omitempty"`
		// Maximum number of applications
		MaxApplicationsPerAccount int `json:"max_applications_per_account,omitempty"`
		// Cooldown peroid after leaving the clan (in seconds)
		LeaveCooldown int `json:"leave_cooldown,omitempty"`
		// Role received after retirement as the clan commander
		LeaderRoleAfterRetirement string `json:"leader_role_after_retirement,omitempty"`
		// Invitation expiry date
		InviteLifetime int `json:"invite_lifetime,omitempty"`
		// Role received after joining the clan
		DefaultAccountRole string `json:"default_account_role,omitempty"`
		// Clan creation cost in gold
		CreateFeeGold int `json:"create_fee_gold,omitempty"`
		// Application expiry date
		ApplicationLifetime int `json:"application_lifetime,omitempty"`
	} `json:"settings,omitempty"`
	// Available emblems
	Emblems struct {
		// Link to 64x64 emblem
		X64 string `json:"x64,omitempty"`
		// Link to 32x32 emblem
		X32 string `json:"x32,omitempty"`
		// Link to 256x256 emblem
		X256 string `json:"x256,omitempty"`
		// Link to 24x24 emblem
		X24 string `json:"x24,omitempty"`
		// Link to 195x195 emblem
		X195 string `json:"x195,omitempty"`
	} `json:"emblems,omitempty"`
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
} 
