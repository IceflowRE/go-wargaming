package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type TanksMastery struct {
	// Date of data update
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Values of these percentiles for each piece of equipment
	Distribution map[string]string `json:"distribution,omitempty"`
} 
