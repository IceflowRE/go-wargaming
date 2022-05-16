package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type GlobalmapSeasonratingneighbors struct {
	// Victory Points to get the next award
	VictoryPointsToNextAward int `json:"victory_points_to_next_award,omitempty"`
	// Victory Points
	VictoryPoints int `json:"victory_points,omitempty"`
	// Date of rating calculation
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
	// Clan tag
	Tag string `json:"tag,omitempty"`
	// Rating changes during previous turn
	RankDelta int `json:"rank_delta,omitempty"`
	// Current rating
	Rank int `json:"rank,omitempty"`
	// Clan name
	Name string `json:"name,omitempty"`
	// Clan color
	Color string `json:"color,omitempty"`
	// Clan ID
	ClanId int `json:"clan_id,omitempty"`
	// Award level
	AwardLevel string `json:"award_level,omitempty"`
} 
