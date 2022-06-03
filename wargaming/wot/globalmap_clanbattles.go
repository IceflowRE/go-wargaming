package wot

import (
	"github.com/IceflowRE/go-wargaming/v2/wargaming/wgnTime"
)

// GlobalmapClanbattlesOptions options.
type GlobalmapClanbattlesOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Language. Default is "ru". Valid values:
	// 
	// "ru" - Russian (by default)
	Language *string
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int
	// Page number. Default is 1. Min value is 1.
	PageNo *int
}

type GlobalmapClanbattles struct {
	// Attack Type: ground, auction, tournament
	AttackType *string `json:"attack_type,omitempty"`
	// Enemy clan ID
	CompetitorId *int `json:"competitor_id,omitempty"`
	// Front ID
	FrontId *string `json:"front_id,omitempty"`
	// Front name
	FrontName *string `json:"front_name,omitempty"`
	// Province ID
	ProvinceId *string `json:"province_id,omitempty"`
	// Province name
	ProvinceName *string `json:"province_name,omitempty"`
	// Battle date and time
	Time *wgnTime.UnixTime `json:"time,omitempty"`
	// Battle type: attack or defense
	Type_ *string `json:"type,omitempty"`
	// Vehicle Tier
	VehicleLevel *int `json:"vehicle_level,omitempty"`
}
