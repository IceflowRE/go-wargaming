package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapClanbattles struct {
	// Vehicle Tier
	VehicleLevel int `json:"vehicle_level,omitempty"`
	// Battle type: attack or defense
	Type_ string `json:"type,omitempty"`
	// Battle date and time
	Time wgnTime.UnixTime `json:"time,omitempty"`
	// Province name
	ProvinceName string `json:"province_name,omitempty"`
	// Province ID
	ProvinceId string `json:"province_id,omitempty"`
	// Front name
	FrontName string `json:"front_name,omitempty"`
	// Front ID
	FrontId string `json:"front_id,omitempty"`
	// Enemy clan ID
	CompetitorId int `json:"competitor_id,omitempty"`
	// Attack Type: ground, auction, tournament
	AttackType string `json:"attack_type,omitempty"`
} 
