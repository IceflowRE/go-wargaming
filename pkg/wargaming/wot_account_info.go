package wargaming

type statisticGeneral struct {
	AvgDamageAssisted          float32 `json:"avg_damage_assisted"`
	AvgDamageAssistedRadio     float32 `json:"avg_damage_assisted_radio"`
	AvgDamageAssistedTrack     float32 `json:"avg_damage_assisted_track"`
	AvgDamagedBlocked          float32 `json:"avg_damaged_blocked"`
	BattleAvgXp                int     `json:"battle_avg_xp"`
	Battles                    int     `json:"battles"`
	BattlesOnStunningVehicles  int     `json:"battles_on_stunning_vehicles"`
	CapturePoints              int     `json:"capture_points"`
	DamageDealt                int     `json:"damage_dealt"`
	DamageReceived             int     `json:"damage_received"`
	DirectHitsReceived         int     `json:"direct_hits_received"`
	Draws                      int     `json:"draws"`
	DroppedCapturePoints       int     `json:"dropped_capture_points"`
	ExplosionHits              int     `json:"explosion_hits"`
	ExplosionHitsReceived      int     `json:"explosion_hits_received"`
	Frags                      int     `json:"frags"`
	Hits                       int     `json:"hits"`
	HitsPercents               int     `json:"hits_percents"`
	Losses                     int     `json:"losses"`
	NoDamageDirectHitsReceived int     `json:"no_damage_direct_hits_received"`
	Piercings                  int     `json:"piercings"`
	PiercingsReceived          int     `json:"piercings_received"`
	Shots                      int     `json:"shots"`
	Spotted                    int     `json:"spotted"`
	StunAssistedDamage         int     `json:"stun_assisted_damage"`
	StunNumber                 int     `json:"stun_number"`
	SurvivedBattles            int     `json:"survived_battles"`
	TankingFactor              float32 `json:"tanking_factor"`
	Wins                       int     `json:"wins"`
	Xp                         int     `json:"xp"`
}

type statisticAdvanced struct {
	statisticGeneral
	MaxDamage       int `json:"max_damage"`
	MaxDamageTankId int `json:"max_damage_tank_id"`
	MaxFrags        int `json:"max_frags"`
	MaxFragsTankId  int `json:"max_frags_tank_id"`
	MaxXp           int `json:"max_xp"`
	MaxXpTankId     int `json:"max_xp_tank_id"`
}

type WotAccountInfo struct {
	AccountId      int      `json:"account_id"`
	ClanId         int      `json:"clan_id"`
	ClientLanguage string   `json:"client_language"`
	CreatedAt      UnixTime `json:"created_at"`
	GlobalRating   int      `json:"global_rating"`
	LastBattleTime UnixTime `json:"last_battle_time"`
	LogoutAt       UnixTime `json:"logout_at"`
	Nickname       string   `json:"nickname"`
	UpdatedAt      UnixTime `json:"updated_at"`
	Private        struct {
		BanInfo          string            `json:"ban_info"`
		BanTime          UnixTime          `json:"ban_time"`
		BattleLifeTime   int               `json:"battle_life_time"`
		Bonds            int               `json:"bonds"`
		Credits          int               `json:"credits"`
		FreeXp           int               `json:"free_xp"`
		Garage           []int             `json:"garage"`
		Gold             int               `json:"gold"`
		IsBoundToPhone   bool              `json:"is_bound_to_phone"`
		IsPremium        bool              `json:"is_premium"`
		PersonalMissions map[string]string `json:"personal_missions"`
		PremiumExpiresAt UnixTime          `json:"premium_expires_at"`
		Boosters         map[string]struct {
			Count          int      `json:"count"`
			State          string   `json:"state"`
			ExpirationTime UnixTime `json:"expirationTime"`
		} `json:"boosters"`
	} `json:"private"`
	GroupedContacts struct {
		Blocked   []int             `json:"blocked"`
		Groups    map[string]string `json:"groups"`
		Ignored   []int             `json:"ignored"`
		Muted     []int             `json:"muted"`
		Ungrouped []int             `json:"ungrouped"`
	} `json:"grouped_contacts"`
	Rented map[string]struct {
		CompensationCredits int      `json:"compensation_credits"`
		CompensationGold    int      `json:"compensation_gold"`
		ExpirationTime      UnixTime `json:"expiration_time"`
		TankId              int      `json:"tank_id"`
	} `json:"rented"`
	Restrictions struct {
		ChatBanTime UnixTime `json:"chat_ban_time"`
	} `json:"restrictions"`
	Statistics struct {
		Frags    map[string]int    `json:"frags"`
		TreesCut int               `json:"trees_cut"`
		All      statisticAdvanced `json:"all"`
		Clan     statisticGeneral  `json:"clan"`
		Company  statisticGeneral  `json:"company"`
		Epic     statisticAdvanced `json:"epic"`
		Fallout  struct {
			statisticAdvanced
			AvatarDamageDealt   int `json:"avatar_damage_dealt"`
			AvatarFrags         int `json:"avatar_frags"`
			MaxDamageWithAvatar int `json:"max_damage_with_avatar"`
			MaxFragsWithAvatar  int `json:"max_frags_with_avatar"`
			MaxWinPoints        int `json:"max_win_points"`
			ResourceAbsorbed    int `json:"resource_absorbed"`
			WinPoints           int `json:"win_points"`
		} `json:"fallout"`
		GlobalmapAbsolute     statisticGeneral  `json:"globalmap_absolute"`
		GlobalmapChampion     statisticGeneral  `json:"globalmap_champion"`
		GlobalmapMiddle       statisticGeneral  `json:"globalmap_middle"`
		Historical            statisticAdvanced `json:"historical"`
		Random                statisticGeneral  `json:"random"`
		Ranked10x10           statisticAdvanced `json:"ranked_10x10"`
		Ranked15x15           statisticAdvanced `json:"ranked_15x15"`
		RankedBattles         statisticAdvanced `json:"ranked_battles"`
		RankedBattlesCurrent  statisticAdvanced `json:"ranked_battles_current"`
		RankedBattlesPrevious statisticAdvanced `json:"ranked_battles_previous"`
		RankedSeason1         statisticAdvanced `json:"ranked_season_1"`
		RankedSeason2         statisticAdvanced `json:"ranked_season_2"`
		RankedSeason3         statisticAdvanced `json:"ranked_season_3"`
		RegularTeam           statisticAdvanced `json:"regular_team"`
		StrongholdDefense     statisticAdvanced `json:"stronghold_defense"`
		StrongholdSkirmish    statisticAdvanced `json:"stronghold_skirmish"`
		Team                  statisticAdvanced `json:"team"`
	} `json:"statistics"`
}

func (client *Client) WotAccountInfo(realm Realm, accountId string, accessToken string, extra string, fields string, language string) ([]*WotAccountInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id":   accountId,
		"access_token": accessToken,
		"extra":        extra,
		"fields":       fields,
		"language":     language,
	}
	var result []*WotAccountInfo
	err := client.sendGetRequest(realm, "/wot/account/info/", reqParam, &result)
	return result, err
}
