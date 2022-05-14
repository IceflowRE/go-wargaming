package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowsSeasonsAccountinfo struct {
	// User ID
	AccountId int `json:"account_id,omitempty"`
	// Ranked Battles seasons
	Seasons struct {
		// Player statistics in Ranked Battles in division of 2 players
		RankDiv2 struct {
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"rank_div2,omitempty"`
		// Player statistics in Ranked Battles in division of 3 players
		RankDiv3 struct {
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"rank_div3,omitempty"`
		// Rank Battles information
		RankInfo struct {
			// Player's best rank in current season
			MaxRank int `json:"max_rank,omitempty"`
			// Player's current rank
			Rank int `json:"rank,omitempty"`
			// Current season status
			Stage int `json:"stage,omitempty"`
			// Number of stars
			Stars int `json:"stars,omitempty"`
			// Rank required to participate in a season
			StartRank int `json:"start_rank,omitempty"`
		} `json:"rank_info,omitempty"`
		// Player statistics in Ranked Battles in solo game
		RankSolo struct {
			// Battles fought
			Battles int `json:"battles,omitempty"`
			// Damage caused
			DamageDealt int `json:"damage_dealt,omitempty"`
			// Draws
			Draws int `json:"draws,omitempty"`
			// Warships destroyed
			Frags int `json:"frags,omitempty"`
			// Defeats
			Losses int `json:"losses,omitempty"`
			// Maximum Damage caused per battle
			MaxDamageDealt int `json:"max_damage_dealt,omitempty"`
			// Maximum number of enemy warships destroyed per battle
			MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			// Maximum number of aircraft destroyed per battle
			MaxPlanesKilled int `json:"max_planes_killed,omitempty"`
			// Maximum Experience Points per battle, including XP earned with Premium Account
			MaxXp int `json:"max_xp,omitempty"`
			// Enemy aircraft destroyed
			PlanesKilled int `json:"planes_killed,omitempty"`
			// Battles survived
			SurvivedBattles int `json:"survived_battles,omitempty"`
			// Victories in battles survived
			SurvivedWins int `json:"survived_wins,omitempty"`
			// Victories
			Wins int `json:"wins,omitempty"`
			// Total Experience Points , including XP earned with Premium Account
			Xp int `json:"xp,omitempty"`
			// Statistics of aircraft used
			Aircraft struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			} `json:"aircraft,omitempty"`
			// Main battery firing statistics
			MainBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"main_battery,omitempty"`
			// Statistics of ramming enemy warships
			Ramming struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
			} `json:"ramming,omitempty"`
			// Secondary armament firing statistics
			SecondBattery struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"second_battery,omitempty"`
			// Statistics of attacking targets with torpedoes
			Torpedoes struct {
				// Warships destroyed
				Frags int `json:"frags,omitempty"`
				// Hits
				Hits int `json:"hits,omitempty"`
				// Maximum number of enemy warships destroyed per battle
				MaxFragsBattle int `json:"max_frags_battle,omitempty"`
				// Shots fired
				Shots int `json:"shots,omitempty"`
			} `json:"torpedoes,omitempty"`
		} `json:"rank_solo,omitempty"`
	} `json:"seasons,omitempty"`
}

// WowsSeasonsAccountinfo Method returns players' statistics in Ranked Battles seasons. Accounts with hidden game profiles are excluded from response. Hidden profiles are listed in the field meta.hidden.
//
// https://developers.wargaming.net/reference/all/wows/seasons/accountinfo
//
// account_id:
//     Player account ID. Maximum limit: 100. Min value is 1.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
// season_id:
//     Season ID. Maximum limit: 100.
func (client *Client) WowsSeasonsAccountinfo(realm Realm, accountId []int, accessToken string, fields []string, language string, seasonId []int) (*WowsSeasonsAccountinfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
		"season_id": utils.SliceIntToString(seasonId, ","),
	}

	var result *WowsSeasonsAccountinfo
	err := client.doGetDataRequest(realm, "/wows/seasons/accountinfo/", reqParam, &result)
	return result, err
}
