package wot

type GlobalmapProvinces struct {
	// Indicates if Repartition of the World is active
	WorldRedivision bool `json:"world_redivision,omitempty"`
	// Relative link to province
	Uri string `json:"uri,omitempty"`
	// Tournament status: STARTED, FINISHED or null
	Status string `json:"status,omitempty"`
	// Server ID
	Server string `json:"server,omitempty"`
	// Round
	RoundNumber int `json:"round_number,omitempty"`
	// Income level from 0 to 11. 0 value means that income was not raised.
	RevenueLevel int `json:"revenue_level,omitempty"`
	// Province name
	ProvinceName string `json:"province_name,omitempty"`
	// Province ID
	ProvinceId string `json:"province_id,omitempty"`
	// Prime Time in UTC
	PrimeTime string `json:"prime_time,omitempty"`
	// Date when province will restore its revenue after ransack
	PillageEndAt string `json:"pillage_end_at,omitempty"`
	// Owning clan ID
	OwnerClanId int `json:"owner_clan_id,omitempty"`
	// List of adjacent provinces' IDs
	Neighbours []string `json:"neighbours,omitempty"`
	// Maximum number of bids
	MaxBets int `json:"max_bets,omitempty"`
	// Last winning bid
	LastWonBet int `json:"last_won_bet,omitempty"`
	// Landing type: auction, tournament or null
	LandingType string `json:"landing_type,omitempty"`
	// Province borders are closed
	IsBordersDisabled bool `json:"is_borders_disabled,omitempty"`
	// Front name
	FrontName string `json:"front_name,omitempty"`
	// Front ID
	FrontId string `json:"front_id,omitempty"`
	// Daily income
	DailyRevenue int `json:"daily_revenue,omitempty"`
	// Current minimum bid
	CurrentMinBet int `json:"current_min_bet,omitempty"`
	// List of IDs of participating clans
	Competitors []int `json:"competitors,omitempty"`
	// Battles start time in UTC
	BattlesStartAt string `json:"battles_start_at,omitempty"`
	// List of IDs of attacking clans
	Attackers []int `json:"attackers,omitempty"`
	// Localized map name
	ArenaName string `json:"arena_name,omitempty"`
	// Map ID
	ArenaId string `json:"arena_id,omitempty"`
	// Current battles
	ActiveBattles struct {
		// Battle start time in UTC
		StartAt string `json:"start_at,omitempty"`
		// Round
		Round int `json:"round,omitempty"`
		// Second challenging clan ID
		ClanB struct {
			// Changes in Elo-rating due to victory
			WinEloDelta int `json:"win_elo_delta,omitempty"`
			// Changes in Elo-rating due to defeat
			LooseEloDelta int `json:"loose_elo_delta,omitempty"`
			// Clan ID
			ClanId int `json:"clan_id,omitempty"`
			// Award
			BattleReward int `json:"battle_reward,omitempty"`
		} `json:"clan_b,omitempty"`
		// First challenging clan ID
		ClanA struct {
			// Changes in Elo-rating due to victory
			WinEloDelta int `json:"win_elo_delta,omitempty"`
			// Changes in Elo-rating due to defeat
			LooseEloDelta int `json:"loose_elo_delta,omitempty"`
			// Clan ID
			ClanId int `json:"clan_id,omitempty"`
			// Award
			BattleReward int `json:"battle_reward,omitempty"`
		} `json:"clan_a,omitempty"`
		// Award
		BattleReward int `json:"battle_reward,omitempty"`
	} `json:"active_battles,omitempty"`
} 
