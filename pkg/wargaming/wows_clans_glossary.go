package wargaming

import (
	"strings"
)

type WowsClansGlossary struct {
	// Available clan positions
	ClansRoles map[string]string `json:"clans_roles,omitempty"`
	// Installation type
	BuildingTypes struct {
		// Installation ID
		BuildingTypeId int `json:"building_type_id,omitempty"`
		// Structure name
		Name string `json:"name,omitempty"`
	} `json:"building_types,omitempty"`
	// Installations
	Buildings struct {
		// The type of the bonus that is provided to the clan members after building the installation. Existing bonus types:
		// 
		// exp_boost—rate of additional XP;
		// members_count—number of clan members on which the clan size is increased;
		// maintenance_discount—rate of reducing cost of servicing ships;
		// purchase_discount—rate of reducing cost of researched ships;
		// dummy—no bonuses.
		BonusType string `json:"bonus_type,omitempty"`
		// The value of the bonus based on this bonus type.
		BonusValue int `json:"bonus_value,omitempty"`
		// Installation ID
		BuildingId int `json:"building_id,omitempty"`
		// Installation ID
		BuildingTypeId int `json:"building_type_id,omitempty"`
		// Building cost in oil
		Cost int `json:"cost,omitempty"`
		// Installation name
		Name string `json:"name,omitempty"`
		// The nation of ships that will get a bonus from this installation. If value = "empty string" or "null", the bonus will be received on ships of all nations.
		ShipNation string `json:"ship_nation,omitempty"`
		// Tier of ships that will get a bonus after building this installation. If value = "null", the bonus will be received on ships of all Tiers.
		ShipTier int `json:"ship_tier,omitempty"`
		// Type of ships that will get a bonus after building this installation. If value = "empty string" or "null", the bonus will be received on ships of all types.
		ShipType string `json:"ship_type,omitempty"`
	} `json:"buildings,omitempty"`
	// Clan settings
	Settings struct {
		// Max number of clan members
		MaxMembersCount int `json:"max_members_count,omitempty"`
	} `json:"settings,omitempty"`
}

// WowsClansGlossary Method returns information on clan entities.
//
// https://developers.wargaming.net/reference/all/wows/clans/glossary
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
func (client *Client) WowsClansGlossary(realm Realm, fields []string, language string) (*WowsClansGlossary, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowsClansGlossary
	err := client.doGetDataRequest(realm, "/wows/clans/glossary/", reqParam, &result)
	return result, err
}
