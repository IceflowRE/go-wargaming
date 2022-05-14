package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotxEncyclopediaVehicles struct {
	// Vehicle description
	Description string `json:"description,omitempty"`
	// Era
	Era string `json:"era,omitempty"`
	// Indicates if the vehicle is Premium vehicle
	IsPremium bool `json:"is_premium,omitempty"`
	// Vehicle name
	Name string `json:"name,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// List of vehicles available for research in form of pairs:
	// 
	// researched vehicle ID
	// cost of research in XP
	NextTanks map[string]string `json:"next_tanks,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Cost in gold
	PriceGold int `json:"price_gold,omitempty"`
	// List of research costs in form of pairs:
	// 
	// parent vehicle ID
	// cost of research in XP
	PricesXp map[string]string `json:"prices_xp,omitempty"`
	// Vehicle short name
	ShortName string `json:"short_name,omitempty"`
	// Vehicle tag
	Tag string `json:"tag,omitempty"`
	// Vehicle ID
	TankId int `json:"tank_id,omitempty"`
	// Tier
	Tier int `json:"tier,omitempty"`
	// Vehicle type
	Type_ string `json:"type,omitempty"`
	// Image links
	Images struct {
		// URL to 160 x 100 px image
		BigIcon string `json:"big_icon,omitempty"`
	} `json:"images,omitempty"`
}

// WotxEncyclopediaVehicles Method returns list of available vehicles.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicles
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
// nation:
//     Nation. Maximum limit: 100.
// tank_id:
//     Vehicle ID. Maximum limit: 100.
// tier:
//     Tier. Maximum limit: 100.
func (client *Client) WotxEncyclopediaVehicles(realm Realm, fields []string, language string, nation []string, tankId []int, tier []int) (*WotxEncyclopediaVehicles, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"nation": strings.Join(nation, ","),
		"tank_id": utils.SliceIntToString(tankId, ","),
		"tier": utils.SliceIntToString(tier, ","),
	}

	var result *WotxEncyclopediaVehicles
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/vehicles/", reqParam, &result)
	return result, err
}
