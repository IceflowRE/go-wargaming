package wot

import (
	"github.com/IceflowRE/go-wargaming/wargaming/wgnTime"
)

// StrongholdActivateclanreserveOptions options.
type StrongholdActivateclanreserveOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	// 
	// "ru" - Russian (by default)
	// "be" - Belarusian 
	// "uk" - Ukrainian 
	// "kk" - Kazakh
	Language *string
}

type StrongholdActivateclanreserve struct {
	// Activation time of a clan Reserve
	ActivatedAt *wgnTime.UnixTime `json:"activated_at,omitempty"`
}
