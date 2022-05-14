package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WotxEncyclopediaVehicleupgrades struct {
	// List of compatible consumables
	Consumables struct {
		// Consumable ID
		ConsumableId int `json:"consumable_id,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Cost in gold
		PriceGold int `json:"price_gold,omitempty"`
	} `json:"consumables,omitempty"`
	// List of compatible equipment
	Equipment struct {
		// Achievement description
		Description string `json:"description,omitempty"`
		// Equipment ID
		EquipmentId int `json:"equipment_id,omitempty"`
		// URL to image
		Image string `json:"image,omitempty"`
		// Vehicle name
		Name string `json:"name,omitempty"`
		// Cost in credits
		PriceCredit int `json:"price_credit,omitempty"`
		// Cost in gold
		PriceGold int `json:"price_gold,omitempty"`
	} `json:"equipment,omitempty"`
}

// WotxEncyclopediaVehicleupgrades Method returns list of vehicle equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/vehicleupgrades
//
// tank_id:
//     Vehicle ID. Maximum limit: 100.
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
func (client *Client) WotxEncyclopediaVehicleupgrades(realm Realm, tankId []int, fields []string, language string) (*WotxEncyclopediaVehicleupgrades, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": utils.SliceIntToString(tankId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotxEncyclopediaVehicleupgrades
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/vehicleupgrades/", reqParam, &result)
	return result, err
}
