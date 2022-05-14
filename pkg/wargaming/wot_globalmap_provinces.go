package wargaming

import (
	"strings"
	"strconv"
)

type WotGlobalmapProvinces struct {
	// Map ID
	ArenaId string `json:"arena_id,omitempty"`
	// Localized map name
	ArenaName string `json:"arena_name,omitempty"`
	// List of IDs of attacking clans
	Attackers []int `json:"attackers,omitempty"`
	// Battles start time in UTC
	BattlesStartAt string `json:"battles_start_at,omitempty"`
	// List of IDs of participating clans
	Competitors []int `json:"competitors,omitempty"`
	// Current minimum bid
	CurrentMinBet int `json:"current_min_bet,omitempty"`
	// Daily income
	DailyRevenue int `json:"daily_revenue,omitempty"`
	// Front ID
	FrontId string `json:"front_id,omitempty"`
	// Front name
	FrontName string `json:"front_name,omitempty"`
	// Province borders are closed
	IsBordersDisabled bool `json:"is_borders_disabled,omitempty"`
	// Landing type: auction, tournament or null
	LandingType string `json:"landing_type,omitempty"`
	// Last winning bid
	LastWonBet int `json:"last_won_bet,omitempty"`
	// Maximum number of bids
	MaxBets int `json:"max_bets,omitempty"`
	// List of adjacent provinces' IDs
	Neighbours []string `json:"neighbours,omitempty"`
	// Owning clan ID
	OwnerClanId int `json:"owner_clan_id,omitempty"`
	// Date when province will restore its revenue after ransack
	PillageEndAt string `json:"pillage_end_at,omitempty"`
	// Prime Time in UTC
	PrimeTime string `json:"prime_time,omitempty"`
	// Province ID
	ProvinceId string `json:"province_id,omitempty"`
	// Province name
	ProvinceName string `json:"province_name,omitempty"`
	// Income level from 0 to 11. 0 value means that income was not raised.
	RevenueLevel int `json:"revenue_level,omitempty"`
	// Round
	RoundNumber int `json:"round_number,omitempty"`
	// Server ID
	Server string `json:"server,omitempty"`
	// Tournament status: STARTED, FINISHED or null
	Status string `json:"status,omitempty"`
	// Relative link to province
	Uri string `json:"uri,omitempty"`
	// Indicates if Repartition of the World is active
	WorldRedivision bool `json:"world_redivision,omitempty"`
	// Current battles
	ActiveBattles struct {
		// Award
		BattleReward int `json:"battle_reward,omitempty"`
		// Round
		Round int `json:"round,omitempty"`
		// Battle start time in UTC
		StartAt string `json:"start_at,omitempty"`
		// First challenging clan ID
		ClanA struct {
			// Award
			BattleReward int `json:"battle_reward,omitempty"`
			// Clan ID
			ClanId int `json:"clan_id,omitempty"`
			// Changes in Elo-rating due to defeat
			LooseEloDelta int `json:"loose_elo_delta,omitempty"`
			// Changes in Elo-rating due to victory
			WinEloDelta int `json:"win_elo_delta,omitempty"`
		} `json:"clan_a,omitempty"`
		// Second challenging clan ID
		ClanB struct {
			// Award
			BattleReward int `json:"battle_reward,omitempty"`
			// Clan ID
			ClanId int `json:"clan_id,omitempty"`
			// Changes in Elo-rating due to defeat
			LooseEloDelta int `json:"loose_elo_delta,omitempty"`
			// Changes in Elo-rating due to victory
			WinEloDelta int `json:"win_elo_delta,omitempty"`
		} `json:"clan_b,omitempty"`
	} `json:"active_battles,omitempty"`
}

// WotGlobalmapProvinces Method returns information about the Global Map provinces.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/provinces
//
// arena_id:
//     Map ID
// daily_revenue_gte:
//     Search for provinces with daily income equal to or more than the value
// daily_revenue_lte:
//     Search for provinces with daily income equal to or less than the value
// front_id:
//     Front ID. To get a front ID, use the Fronts method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// landing_type:
//     Search for provinces by landing type. Valid values:
//     
//     "null" &mdash; auctions disabled 
//     "auction" &mdash; auction 
//     "tournament" &mdash; landing tournament
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// order_by:
//     Sorting. Valid values:
//     
//     "province_id" &mdash; by province name 
//     "-province_id" &mdash; by province name in reverse order 
//     "daily_revenue" &mdash; by province income 
//     "-daily_revenue" &mdash; by province income in reverse order 
//     "prime_hour" &mdash; by Prime Time 
//     "-prime_hour" &mdash; by Prime Time in reverse order
// page_no:
//     Page number. Default is 1. Min value is 1.
// prime_hour:
//     Search for provinces with the value of Prime Time start hour. Values available: from 0 to 23. Maximum value: 23.
// province_id:
//     Filter by the list of province IDs. Maximum limit: 100.
func (client *Client) WotGlobalmapProvinces(realm Realm, arenaId string, dailyRevenueGte int, dailyRevenueLte int, frontId string, fields []string, landingType string, language string, limit int, orderBy string, pageNo int, primeHour int, provinceId []string) (*WotGlobalmapProvinces, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"arena_id": arenaId,
		"daily_revenue_gte": strconv.Itoa(dailyRevenueGte),
		"daily_revenue_lte": strconv.Itoa(dailyRevenueLte),
		"front_id": frontId,
		"fields": strings.Join(fields, ","),
		"landing_type": landingType,
		"language": language,
		"limit": strconv.Itoa(limit),
		"order_by": orderBy,
		"page_no": strconv.Itoa(pageNo),
		"prime_hour": strconv.Itoa(primeHour),
		"province_id": strings.Join(provinceId, ","),
	}

	var result *WotGlobalmapProvinces
	err := client.doGetDataRequest(realm, "/wot/globalmap/provinces/", reqParam, &result)
	return result, err
}
