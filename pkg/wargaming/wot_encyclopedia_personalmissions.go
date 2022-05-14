package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotEncyclopediaPersonalmissions struct {
	// Campaign ID
	CampaignId int `json:"campaign_id,omitempty"`
	// Campaign description
	Description string `json:"description,omitempty"`
	// Campaign name
	Name string `json:"name,omitempty"`
	// Campaign operations
	Operations struct {
		// Operation description
		Description string `json:"description,omitempty"`
		// Link to an operation image
		Image string `json:"image,omitempty"`
		// Number of missions in the branch
		MissionsInSet int `json:"missions_in_set,omitempty"`
		// Operation name
		Name string `json:"name,omitempty"`
		// Next operation ID
		NextId int `json:"next_id,omitempty"`
		// Operation ID
		OperationId int `json:"operation_id,omitempty"`
		// Number of mission branches of the operation
		SetsCount int `json:"sets_count,omitempty"`
		// Number of branches until the next operation
		SetsToNext int `json:"sets_to_next,omitempty"`
		// Operation missions
		Missions struct {
			// Mission description
			Description string `json:"description,omitempty"`
			// Hints
			Hint string `json:"hint,omitempty"`
			// Maximum vehicle Tier
			MaxLevel int `json:"max_level,omitempty"`
			// Minimum vehicle Tier
			MinLevel int `json:"min_level,omitempty"`
			// Mission ID
			MissionId int `json:"mission_id,omitempty"`
			// Mission name
			Name string `json:"name,omitempty"`
			// Mission branch ID
			SetId int `json:"set_id,omitempty"`
			// Mission tags
			Tags []string `json:"tags,omitempty"`
			// Rewards grouped by mission conditions
			Rewards struct {
				// Bunks in Barracks
				Berths int `json:"berths,omitempty"`
				// Mission conditions
				Conditions string `json:"conditions,omitempty"`
				// Credits
				Credits int `json:"credits,omitempty"`
				// Free Experience
				FreeXp int `json:"free_xp,omitempty"`
				// List of equipment or consumables in the following format: Id–number of items
				Items map[string]string `json:"items,omitempty"`
				// Days of Premium Account
				Premium int `json:"premium,omitempty"`
				// Slots
				Slots int `json:"slots,omitempty"`
				// Tokens
				Tokens int `json:"tokens,omitempty"`
			} `json:"rewards,omitempty"`
		} `json:"missions,omitempty"`
		// Reward for operation
		Reward struct {
			// Slots
			Slots int `json:"slots,omitempty"`
			// Vehicle IDs
			Tanks []int `json:"tanks,omitempty"`
		} `json:"reward,omitempty"`
	} `json:"operations,omitempty"`
}

// WotEncyclopediaPersonalmissions Method returns details on Personal Missions on the basis of specified campaign IDs, operation IDs, mission branch and tag IDs.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/personalmissions
//
// campaign_id:
//     Campaign ID. Maximum limit: 100.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
// operation_id:
//     Operation ID. Maximum limit: 100.
// set_id:
//     Mission branch ID. Maximum limit: 100.
// tag:
//     Mission tag. Maximum limit: 100.
func (client *Client) WotEncyclopediaPersonalmissions(realm Realm, campaignId []int, fields []string, language string, operationId []int, setId []int, tag []string) (*WotEncyclopediaPersonalmissions, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"campaign_id": utils.SliceIntToString(campaignId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"operation_id": utils.SliceIntToString(operationId, ","),
		"set_id": utils.SliceIntToString(setId, ","),
		"tag": strings.Join(tag, ","),
	}

	var result *WotEncyclopediaPersonalmissions
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/personalmissions/", reqParam, &result)
	return result, err
}
