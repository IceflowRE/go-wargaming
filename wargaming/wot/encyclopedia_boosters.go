// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type EncyclopediaBoostersOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Localization language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "ru" - Русский
	// "pl" - Polski
	// "de" - Deutsch
	// "fr" - Français
	// "es" - Español
	// "zh-cn" - 简体中文
	// "zh-tw" - 繁體中文
	// "tr" - Türkçe
	// "cs" - Čeština
	// "th" - ไทย
	// "vi" - Tiếng Việt
	// "ko" - 한국어
	Language *string `json:"language,omitempty"`
}

type EncyclopediaBoosters struct {
	// Personal Reserve ID
	BoosterId *int `json:"booster_id,omitempty"`
	// Personal Reserve description
	Description *string `json:"description,omitempty"`
	// Personal Reserve expiration time in UTC
	ExpiresAt *wgnTime.UnixTime `json:"expires_at,omitempty"`
	// Personal Reserve image
	Images *struct {
		// URL to large image
		Large *string `json:"large,omitempty"`
		// URL to small image
		Small *string `json:"small,omitempty"`
	} `json:"images,omitempty"`
	// Personal Reserve auto activation flag
	IsAuto *bool `json:"is_auto,omitempty"`
	// Personal Reserve duration in seconds
	Lifetime *int `json:"lifetime,omitempty"`
	// Personal Reserve name
	Name *string `json:"name,omitempty"`
	// Cost in credits
	PriceCredit *int `json:"price_credit,omitempty"`
	// Price in gold
	PriceGold *int `json:"price_gold,omitempty"`
	// Resource affected by Personal Reserve. Valid values: credits, experience, crew_experience, free_experience.
	Resource *string `json:"resource,omitempty"`
}
