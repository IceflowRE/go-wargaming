// Auto generated file!

package wot

type GlobalmapClanprovincesOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string `json:"access_token,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "fr" - French
	// "es" - Spanish
	// "pl" - Polish
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
}

type GlobalmapClanprovinces struct {
	// Map ID
	ArenaId *string `json:"arena_id,omitempty"`
	// Localized map name
	ArenaName *string `json:"arena_name,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Daily income
	DailyRevenue *int `json:"daily_revenue,omitempty"`
	// Front ID
	FrontId *string `json:"front_id,omitempty"`
	// Front name
	FrontName *string `json:"front_name,omitempty"`
	// Indicates the landing type of a province
	LandingType *string `json:"landing_type,omitempty"`
	// Maximum vehicle Tier in a Front
	MaxVehicleLevel *int `json:"max_vehicle_level,omitempty"`
	// Date when province will restore its revenue after ransack
	PillageEndAt *string `json:"pillage_end_at,omitempty"`
	// Prime Time in UTC
	PrimeTime *string `json:"prime_time,omitempty"`
	// Restricted province information
	Private *struct {
		// Indicates availability of connection to the Headquarters
		HqConnected *bool `json:"hq_connected,omitempty"`
		// Province income limit. Valid values:
		//
		//
		// False-this province income is not blocked due to reaching province income limit. Though, it might be blocked by event rules.
		//
		//
		// True-this province income is blocked, as province income limit has been reached.
		//
		IsRevenueLimitExceeded *bool `json:"is_revenue_limit_exceeded,omitempty"`
	} `json:"private,omitempty"`
	// Province ID
	ProvinceId *string `json:"province_id,omitempty"`
	// Province name
	ProvinceName *string `json:"province_name,omitempty"`
	// Income level from 0 to 11. 0 value means that income was not raised.
	RevenueLevel *int `json:"revenue_level,omitempty"`
	// Province owned (number of turns)
	TurnsOwned *int `json:"turns_owned,omitempty"`
}
