package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapInfo struct {
	// Map status: active, frozen, turn_calculation_in_progress
	State string `json:"state,omitempty"`
	// Creation time of the last calculated turn in UTC
	LastTurnCreatedAt wgnTime.UnixTime `json:"last_turn_created_at,omitempty"`
	// Calculation time of the last turn in UTC
	LastTurnCalculatedAt wgnTime.UnixTime `json:"last_turn_calculated_at,omitempty"`
	// Number of last calculated turn
	LastTurn int `json:"last_turn,omitempty"`
} 
