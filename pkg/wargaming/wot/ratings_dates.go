package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type RatingsDates struct {
	// Dates with available ratings
	Dates []wgnTime.UnixTime `json:"dates,omitempty"`
} 
