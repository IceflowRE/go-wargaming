// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type GlobalmapEventratingOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Clans limit. Default is 10. Min value is 1. Maximum value: 100.
	Limit *int `json:"limit,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
}

type GlobalmapEventrating struct {
	// Award level
	AwardLevel *string `json:"award_level,omitempty"`
	// Battle Fame Points
	BattleFamePoints *int `json:"battle_fame_points,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan color
	Color *string `json:"color,omitempty"`
	// Amount of Fame Points needed to improve personal award
	FamePointsToImproveAward *int `json:"fame_points_to_improve_award,omitempty"`
	// Clan name
	Name *string `json:"name,omitempty"`
	// Current rating
	Rank *int `json:"rank,omitempty"`
	// Rating changes during previous turn
	RankDelta *int `json:"rank_delta,omitempty"`
	// Clan tag
	Tag *string `json:"tag,omitempty"`
	// Fame Points for completing a clan task
	TaskFamePoints *int `json:"task_fame_points,omitempty"`
	// Total Fame Points
	TotalFamePoints *int `json:"total_fame_points,omitempty"`
	// Date of rating calculation
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
}
