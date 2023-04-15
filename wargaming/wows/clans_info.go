// Auto generated file!

package wows

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClansInfoOptions struct {
	// Extra fields that will be added to the response. Valid values:
	//
	// "members"
	Extra []string `json:"extra,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "cs" - Čeština
	// "de" - Deutsch
	// "en" - English (by default)
	// "es" - Español
	// "fr" - Français
	// "ja" - 日本語
	// "pl" - Polski
	// "ru" - Русский
	// "th" - ไทย
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "zh-cn" - 中文
	// "pt-br" - Português do Brasil
	// "es-mx" - Español (México)
	Language *string `json:"language,omitempty"`
}

type ClansInfo struct {
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan creation date
	CreatedAt *wgnTime.UnixTime `json:"created_at,omitempty"`
	// Clan creator ID
	CreatorId *int `json:"creator_id,omitempty"`
	// Clan creator's name
	CreatorName *string `json:"creator_name,omitempty"`
	// Clan description
	Description *string `json:"description,omitempty"`
	// Clan has been deleted. The deleted clan data contains valid values for the following fields only: clan_id, is_clan_disbanded, updated_at.
	IsClanDisbanded *bool `json:"is_clan_disbanded,omitempty"`
	// Commander ID
	LeaderId *int `json:"leader_id,omitempty"`
	// Commander's name
	LeaderName *string `json:"leader_name,omitempty"`
	// Clan members.
	// An extra field.
	Members *struct {
		// User ID
		AccountId *int `json:"account_id,omitempty"`
		// Player name
		AccountName *string `json:"account_name,omitempty"`
		// Date when player joined clan
		JoinedAt *wgnTime.UnixTime `json:"joined_at,omitempty"`
		// Technical position name
		Role *string `json:"role,omitempty"`
	} `json:"members,omitempty"`
	// Number of clan members
	MembersCount *int `json:"members_count,omitempty"`
	// List of clan players' IDs
	MembersIds []int `json:"members_ids,omitempty"`
	// Clan name
	Name *string `json:"name,omitempty"`
	// Old clan name
	OldName *string `json:"old_name,omitempty"`
	// Old clan tag
	OldTag *string `json:"old_tag,omitempty"`
	// Time (UTC) when clan name was changed
	RenamedAt *wgnTime.UnixTime `json:"renamed_at,omitempty"`
	// Clan tag
	Tag *string `json:"tag,omitempty"`
	// Time when clan details were updated
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
