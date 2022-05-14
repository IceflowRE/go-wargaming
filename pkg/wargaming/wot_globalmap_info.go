package wargaming

import (
	"strings"
)

type WotGlobalmapInfo struct {
	// Number of last calculated turn
	LastTurn int `json:"last_turn,omitempty"`
	// Calculation time of the last turn in UTC
	LastTurnCalculatedAt UnixTime `json:"last_turn_calculated_at,omitempty"`
	// Creation time of the last calculated turn in UTC
	LastTurnCreatedAt UnixTime `json:"last_turn_created_at,omitempty"`
	// Map status: active, frozen, turn_calculation_in_progress
	State string `json:"state,omitempty"`
}

// WotGlobalmapInfo Method returns general information about the Global Map.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/info
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
func (client *Client) WotGlobalmapInfo(realm Realm, fields []string) (*WotGlobalmapInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
	}

	var result *WotGlobalmapInfo
	err := client.doGetDataRequest(realm, "/wot/globalmap/info/", reqParam, &result)
	return result, err
}
