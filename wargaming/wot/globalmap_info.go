// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type GlobalmapInfoOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
}

type GlobalmapInfo struct {
	// Number of last calculated turn
	LastTurn *int `json:"last_turn,omitempty"`
	// Calculation time of the last turn in UTC
	LastTurnCalculatedAt *wgnTime.UnixTime `json:"last_turn_calculated_at,omitempty"`
	// Creation time of the last calculated turn in UTC
	LastTurnCreatedAt *wgnTime.UnixTime `json:"last_turn_created_at,omitempty"`
	// Map status: active, frozen, turn_calculation_in_progress
	State *string `json:"state,omitempty"`
}
