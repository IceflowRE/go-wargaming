// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type GlobalmapSeasonratingneighborsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Neighbors limit. Default is 10. Min value is 1. Maximum value: 99.
	Limit *int `json:"limit,omitempty"`
}

type GlobalmapSeasonratingneighbors struct {
	// Award level
	AwardLevel *string `json:"award_level,omitempty"`
	// Clan ID
	ClanId *int `json:"clan_id,omitempty"`
	// Clan color
	Color *string `json:"color,omitempty"`
	// Clan name
	Name *string `json:"name,omitempty"`
	// Current rating
	Rank *int `json:"rank,omitempty"`
	// Rating changes during previous turn
	RankDelta *int `json:"rank_delta,omitempty"`
	// Clan tag
	Tag *string `json:"tag,omitempty"`
	// Date of rating calculation
	UpdatedAt *wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Victory Points
	VictoryPoints *int `json:"victory_points,omitempty"`
	// Victory Points to get the next award
	VictoryPointsToNextAward *int `json:"victory_points_to_next_award,omitempty"`
}
