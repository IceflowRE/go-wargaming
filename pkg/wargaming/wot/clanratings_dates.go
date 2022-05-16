package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type ClanratingsDates struct {
	// List of dates with available ratings
	Dates []wgnTime.UnixTime `json:"dates,omitempty"`
} 
