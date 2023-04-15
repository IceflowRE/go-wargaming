// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type StrongholdActivateclanreserveOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "pl" - Polish
	// "fr" - French
	// "es" - Spanish
	// "cs" - Czech
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
}

type StrongholdActivateclanreserve struct {
	// Activation time of a clan Reserve
	ActivatedAt *wgnTime.UnixTime `json:"activated_at,omitempty"`
}
