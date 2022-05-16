package wot

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type EncyclopediaBoosters struct {
	// Resource affected by Personal Reserve. Valid values: credits, experience, crew_experience, free_experience.
	Resource string `json:"resource,omitempty"`
	// Price in gold
	PriceGold int `json:"price_gold,omitempty"`
	// Cost in credits
	PriceCredit int `json:"price_credit,omitempty"`
	// Personal Reserve name
	Name string `json:"name,omitempty"`
	// Personal Reserve duration in seconds
	Lifetime int `json:"lifetime,omitempty"`
	// Personal Reserve auto activation flag
	IsAuto bool `json:"is_auto,omitempty"`
	// Personal Reserve image
	Images struct {
		// URL to small image
		Small string `json:"small,omitempty"`
		// URL to large image
		Large string `json:"large,omitempty"`
	} `json:"images,omitempty"`
	// Personal Reserve expiration time in UTC
	ExpiresAt wgnTime.UnixTime `json:"expires_at,omitempty"`
	// Personal Reserve description
	Description string `json:"description,omitempty"`
	// Personal Reserve ID
	BoosterId int `json:"booster_id,omitempty"`
} 
