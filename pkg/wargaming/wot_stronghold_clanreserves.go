package wargaming

import (
	"strings"
)

type WotStrongholdClanreserves struct {
	// Reserve efficiency type
	BonusType string `json:"bonus_type,omitempty"`
	// Indicates if the Reserve is a One-Time-Effect Reserve
	Disposable bool `json:"disposable,omitempty"`
	// URL to icon
	Icon string `json:"icon,omitempty"`
	// Reserve name
	Name string `json:"name,omitempty"`
	// Reserve type
	Type_ string `json:"type,omitempty"`
	// Available clan Reserves and their status
	InStock struct {
		// Duration of clan Reserves of each level
		ActionTime int `json:"action_time,omitempty"`
		// Activation time of clan Reserves of each level
		ActivatedAt int `json:"activated_at,omitempty"`
		// Expiration time of clan Reserves of each level
		ActiveTill int `json:"active_till,omitempty"`
		// Number of clan Reserves of each level
		Amount int `json:"amount,omitempty"`
		// Level of available clan Reserves
		Level int `json:"level,omitempty"`
		// Status of clan Reserves of each level
		Status string `json:"status,omitempty"`
		// Indicates if the Reserve is only for Tier X vehicles
		XLevelOnly bool `json:"x_level_only,omitempty"`
		// Reserve efficiencies
		BonusValues struct {
			// Battle type
			BattleType string `json:"battle_type,omitempty"`
			// Reserve efficiency for each battle type
			Value float32 `json:"value,omitempty"`
		} `json:"bonus_values,omitempty"`
	} `json:"in_stock,omitempty"`
}

// WotStrongholdClanreserves Method returns information about available Reserves and their current status.
//
// https://developers.wargaming.net/reference/all/wot/stronghold/clanreserves
//
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
//     "be" &mdash; Belarusian 
//     "uk" &mdash; Ukrainian 
//     "kk" &mdash; Kazakh
func (client *Client) WotStrongholdClanreserves(realm Realm, accessToken string, fields []string, language string) (*WotStrongholdClanreserves, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WotStrongholdClanreserves
	err := client.doGetDataRequest(realm, "/wot/stronghold/clanreserves/", reqParam, &result)
	return result, err
}
