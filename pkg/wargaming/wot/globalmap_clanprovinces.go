package wot

type GlobalmapClanprovinces struct {
	// Province owned (number of turns)
	TurnsOwned int `json:"turns_owned,omitempty"`
	// Income level from 0 to 11. 0 value means that income was not raised.
	RevenueLevel int `json:"revenue_level,omitempty"`
	// Province name
	ProvinceName string `json:"province_name,omitempty"`
	// Province ID
	ProvinceId string `json:"province_id,omitempty"`
	// Restricted province information
	Private struct {
		// Province income limit. Valid values:
		// 
		// 
		// False—this province income is not blocked due to reaching province income limit. Though, it might be blocked by event rules.
		// 
		// 
		// True—this province income is blocked, as province income limit has been reached.
		IsRevenueLimitExceeded bool `json:"is_revenue_limit_exceeded,omitempty"`
		// Indicates availability of connection to the Headquarters
		HqConnected bool `json:"hq_connected,omitempty"`
	} `json:"private,omitempty"`
	// Prime Time in UTC
	PrimeTime string `json:"prime_time,omitempty"`
	// Date when province will restore its revenue after ransack
	PillageEndAt string `json:"pillage_end_at,omitempty"`
	// Maximum vehicle Tier in a Front
	MaxVehicleLevel int `json:"max_vehicle_level,omitempty"`
	// Indicates the landing type of a province
	LandingType string `json:"landing_type,omitempty"`
	// Front name
	FrontName string `json:"front_name,omitempty"`
	// Front ID
	FrontId string `json:"front_id,omitempty"`
	// Daily income
	DailyRevenue int `json:"daily_revenue,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Localized map name
	ArenaName string `json:"arena_name,omitempty"`
	// Map ID
	ArenaId string `json:"arena_id,omitempty"`
} 
