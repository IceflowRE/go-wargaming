package wot

type GlobalmapSeasons struct {
	// Status
	Status string `json:"status,omitempty"`
	// Start time
	Start string `json:"start,omitempty"`
	// Season name
	SeasonName string `json:"season_name,omitempty"`
	// Season ID
	SeasonId string `json:"season_id,omitempty"`
	// Fronts. Only season fronts are presented in response.
	Fronts struct {
		// Link to Front
		Url string `json:"url,omitempty"`
		// Front name
		FrontName string `json:"front_name,omitempty"`
		// Front ID
		FrontId string `json:"front_id,omitempty"`
	} `json:"fronts,omitempty"`
	// Finishing time
	End string `json:"end,omitempty"`
} 
