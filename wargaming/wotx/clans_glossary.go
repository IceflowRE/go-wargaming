// Auto generated file!

package wotx

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
	// "tr" - Türkçe
	Language *string `json:"language,omitempty"`
}

type ClansGlossary struct {
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
	// Available emblems
	Emblems *struct {
		// Link to 195x195 emblem
		X195 *string `json:"x195,omitempty"`
		// Link to 24x24 emblem
		X24 *string `json:"x24,omitempty"`
		// Link to 256x256 emblem
		X256 *string `json:"x256,omitempty"`
		// Link to 32x32 emblem
		X32 *string `json:"x32,omitempty"`
		// Link to 64x64 emblem
		X64 *string `json:"x64,omitempty"`
	} `json:"emblems,omitempty"`
	// Clan settings
	Settings *struct {
		// Application expiry date
		ApplicationLifetime *int `json:"application_lifetime,omitempty"`
		// Clan creation cost in gold
		CreateFeeGold *int `json:"create_fee_gold,omitempty"`
		// Role received after joining the clan
		DefaultAccountRole *string `json:"default_account_role,omitempty"`
		// Invitation expiry date
		InviteLifetime *int `json:"invite_lifetime,omitempty"`
		// Role received after retirement as the clan commander
		LeaderRoleAfterRetirement *string `json:"leader_role_after_retirement,omitempty"`
		// Cooldown peroid after leaving the clan (in seconds)
		LeaveCooldown *int `json:"leave_cooldown,omitempty"`
		// Maximum number of applications
		MaxApplicationsPerAccount *int `json:"max_applications_per_account,omitempty"`
		// Maximum number of invitations
		MaxInvitesPerClan *int `json:"max_invites_per_clan,omitempty"`
		// Max number of clan members
		MaxMembersCount *int `json:"max_members_count,omitempty"`
		// Minimum battles required to create clan
		MinBattlesCount *int `json:"min_battles_count,omitempty"`
		// Cooldown for clan renaming (in seconds)
		RenameCooldown *int `json:"rename_cooldown,omitempty"`
		// Clan renaming cost in gold
		RenameFeeGold *int `json:"rename_fee_gold,omitempty"`
	} `json:"settings,omitempty"`
}
