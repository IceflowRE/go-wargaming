package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WotbTournamentsTables struct {
	// ID of a default clan emblem
	ClanEmblemPresetId int `json:"clan_emblem_preset_id,omitempty"`
	// Team clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Name of a team's clan
	ClanLabel string `json:"clan_label,omitempty"`
	// ID of a team's group
	GroupId int `json:"group_id,omitempty"`
	// Sequence number of a group in a stage
	GroupOrder int `json:"group_order,omitempty"`
	// Number of matches played by a team
	MatchesPlayed int `json:"matches_played,omitempty"`
	// Team's place
	Position int `json:"position,omitempty"`
	// Number of the tour in which a team exited the tournament; relevant only for Single Elimination (SE) and Double Elimination (DE) tournament brackets
	Round int `json:"round,omitempty"`
	// Stage ID
	StageId int `json:"stage_id,omitempty"`
	// Team ID
	TeamId int `json:"team_id,omitempty"`
	// Points earned by a team; relevant only for "Round Robin" tournament brackets
	TeamPoints int `json:"team_points,omitempty"`
	// Team name
	Title string `json:"title,omitempty"`
	// Tournament ID
	TournamentId int `json:"tournament_id,omitempty"`
}

// WotbTournamentsTables Method returns tournament brackets.
//
// https://developers.wargaming.net/reference/all/wotb/tournaments/tables
//
// stage_id:
//     Stage ID that can be retrieved from the Tournaments Stages method.
// tournament_id:
//     Tournament ID that can be retrieved from the Tournaments list method.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// group_id:
//     Group ID. Maximum limit: 10.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Русский (by default)
// limit:
//     Number of returned entries. Default is 10. Min value is 1. Maximum value: 25.
// page_no:
//     Result page number. Default is 1. Min value is 1.
func (client *Client) WotbTournamentsTables(realm Realm, stageId int, tournamentId int, fields []string, groupId []int, language string, limit int, pageNo int) (*WotbTournamentsTables, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"stage_id": strconv.Itoa(stageId),
		"tournament_id": strconv.Itoa(tournamentId),
		"fields": strings.Join(fields, ","),
		"group_id": utils.SliceIntToString(groupId, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
	}

	var result *WotbTournamentsTables
	err := client.doGetDataRequest(realm, "/wotb/tournaments/tables/", reqParam, &result)
	return result, err
}
