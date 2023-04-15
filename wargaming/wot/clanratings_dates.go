// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type ClanratingsDatesOptions struct {
	// Number of returned entries (fewer can be returned, but not more than 365). If the limit sent exceeds 365, a limit of 7 is applied (by default).
	Limit *int `json:"limit,omitempty"`
}

type ClanratingsDates struct {
	// List of dates with available ratings
	Dates []wgnTime.UnixTime `json:"dates,omitempty"`
}
