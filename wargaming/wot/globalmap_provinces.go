// Auto generated file!

package wot

type GlobalmapProvincesOptions struct {
	// Map ID
	ArenaId *string `json:"arena_id,omitempty"`
	// Search for provinces with daily income equal to or more than the value
	DailyRevenueGte *int `json:"daily_revenue_gte,omitempty"`
	// Search for provinces with daily income equal to or less than the value
	DailyRevenueLte *int `json:"daily_revenue_lte,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Search for provinces by landing type. Valid values:
	//
	// "null" - auctions disabled
	// "auction" - auction
	// "tournament" - landing tournament
	LandingType *string `json:"landing_type,omitempty"`
	// Language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "fr" - French
	// "es" - Spanish
	// "pl" - Polish
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
	// Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
	Limit *int `json:"limit,omitempty"`
	// Sorting. Valid values:
	//
	// "province_id" - by province name
	// "-province_id" - by province name in reverse order
	// "daily_revenue" - by province income
	// "-daily_revenue" - by province income in reverse order
	// "prime_hour" - by Prime Time
	// "-prime_hour" - by Prime Time in reverse order
	OrderBy *string `json:"order_by,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
	// Search for provinces with the value of Prime Time start hour. Values available: from 0 to 23. Maximum value: 23.
	PrimeHour *int `json:"prime_hour,omitempty"`
	// Filter by the list of province IDs. Maximum limit: 100.
	ProvinceId []string `json:"province_id,omitempty"`
}

type GlobalmapProvinces struct {
	// Current battles
	ActiveBattles *struct {
		// Award
		BattleReward *int `json:"battle_reward,omitempty"`
		// First challenging clan ID
		ClanA *struct {
			// Award
			BattleReward *int `json:"battle_reward,omitempty"`
			// Clan ID
			ClanId *int `json:"clan_id,omitempty"`
			// Changes in Elo-rating due to defeat
			LooseEloDelta *int `json:"loose_elo_delta,omitempty"`
			// Changes in Elo-rating due to victory
			WinEloDelta *int `json:"win_elo_delta,omitempty"`
		} `json:"clan_a,omitempty"`
		// Second challenging clan ID
		ClanB *struct {
			// Award
			BattleReward *int `json:"battle_reward,omitempty"`
			// Clan ID
			ClanId *int `json:"clan_id,omitempty"`
			// Changes in Elo-rating due to defeat
			LooseEloDelta *int `json:"loose_elo_delta,omitempty"`
			// Changes in Elo-rating due to victory
			WinEloDelta *int `json:"win_elo_delta,omitempty"`
		} `json:"clan_b,omitempty"`
		// Round
		Round *int `json:"round,omitempty"`
		// Battle start time in UTC
		StartAt *string `json:"start_at,omitempty"`
	} `json:"active_battles,omitempty"`
	// Map ID
	ArenaId *string `json:"arena_id,omitempty"`
	// Localized map name
	ArenaName *string `json:"arena_name,omitempty"`
	// List of IDs of attacking clans
	Attackers []int `json:"attackers,omitempty"`
	// Battles start time in UTC
	BattlesStartAt *string `json:"battles_start_at,omitempty"`
	// List of IDs of participating clans
	Competitors []int `json:"competitors,omitempty"`
	// Current minimum bid
	CurrentMinBet *int `json:"current_min_bet,omitempty"`
	// Daily income
	DailyRevenue *int `json:"daily_revenue,omitempty"`
	// Front ID
	FrontId *string `json:"front_id,omitempty"`
	// Front name
	FrontName *string `json:"front_name,omitempty"`
	// Province borders are closed
	IsBordersDisabled *bool `json:"is_borders_disabled,omitempty"`
	// Landing type: auction, tournament or null
	LandingType *string `json:"landing_type,omitempty"`
	// Last winning bid
	LastWonBet *int `json:"last_won_bet,omitempty"`
	// Maximum number of bids
	MaxBets *int `json:"max_bets,omitempty"`
	// List of adjacent provinces' IDs
	Neighbours []string `json:"neighbours,omitempty"`
	// Owning clan ID
	OwnerClanId *int `json:"owner_clan_id,omitempty"`
	// Date when province will restore its revenue after ransack
	PillageEndAt *string `json:"pillage_end_at,omitempty"`
	// Prime Time in UTC
	PrimeTime *string `json:"prime_time,omitempty"`
	// Province ID
	ProvinceId *string `json:"province_id,omitempty"`
	// Province name
	ProvinceName *string `json:"province_name,omitempty"`
	// Income level from 0 to 11. 0 value means that income was not raised.
	RevenueLevel *int `json:"revenue_level,omitempty"`
	// Round
	RoundNumber *int `json:"round_number,omitempty"`
	// Server ID
	Server *string `json:"server,omitempty"`
	// Tournament status: STARTED, FINISHED or null
	Status *string `json:"status,omitempty"`
	// Relative link to province
	Uri *string `json:"uri,omitempty"`
	// Indicates if Repartition of the World is active
	WorldRedivision *bool `json:"world_redivision,omitempty"`
}
