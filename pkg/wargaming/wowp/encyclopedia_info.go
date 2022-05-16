package wowp

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
)

type EncyclopediaInfo struct {
	// Date when Encyclopedia data were updated
	UpdatedAt wgnTime.UnixTime `json:"updated_at,omitempty"`
} 
