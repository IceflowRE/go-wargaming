package wargaming

import (
	"strings"
	"strconv"
)

type WotbTournamentsStages struct {
	// Number of battles in match
	BattleLimit int `json:"battle_limit,omitempty"`
	// Stage description
	Description string `json:"description,omitempty"`
	// Stage end day and time
	EndAt UnixTime `json:"end_at,omitempty"`
	// Number of groups in the stage
	GroupsCount int `json:"groups_count,omitempty"`
	// The allowed maximum vehicle tier for each player in the team
	MaxTier int `json:"max_tier,omitempty"`
	// The allowed minimum vehicle tier for each player in the team
	MinTier int `json:"min_tier,omitempty"`
	// List of stage tours (numbers of tours)
	Rounds []int `json:"rounds,omitempty"`
	// Number of tours in the stage
	RoundsCount int `json:"rounds_count,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Stage start day and time
	StartAt UnixTime `json:"start_at,omitempty"`
	// Stage state. Valid values:
	// 
	// 
	// draft — stage created as draft
	// 
	// 
	// groups_ready —  stage has completed groups
	// 
	// 
	// schedule_ready — stage has a finalized schedule
	// 
	// 
	// complete — stage completed
	State string `json:"state,omitempty"`
	// Stage name
	Title string `json:"title,omitempty"`
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
	// Bracket type of the stage. Valid values:
	// 
	// 
	// RR — round robin
	// 
	// 
	// SE — single elimination
	// 
	// 
	// DE — double elimination
	Type_ string `json:"type,omitempty"`
	// Number of victories to win the match
	VictoryLimit int `json:"victory_limit,omitempty"`
	// Groups info for current stage
	Groups struct {
		// Group ID
		GroupId int `json:"group_id,omitempty"`
		// Group order number
		GroupOrder int `json:"group_order,omitempty"`
	} `json:"groups,omitempty"`
}

// WotbTournamentsStages Method returns list of stages in specified tournament.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/stages
//
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
func (client *Client) WotbTournamentsStages(realm Realm, tournamentId int, fields []string, language string, limit int, pageNo int) (*WotbTournamentsStages, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tournament_id": strconv.Itoa(tournamentId),
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *WotbTournamentsStages
	err := client.doGetDataRequest(realm, "/wotb/tournaments/stages/", reqParam, &result)
	return result, err
}
