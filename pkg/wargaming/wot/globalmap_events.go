package wot

type GlobalmapEvents struct {
	// Status
	Status string `json:"status,omitempty"`
	// Start time
	Start string `json:"start,omitempty"`
	// Fronts. Only event fronts are presented in response.
	Fronts struct {
		// Link to Front
		Url string `json:"url,omitempty"`
		// Front name
		FrontName string `json:"front_name,omitempty"`
		// Front ID
		FrontId string `json:"front_id,omitempty"`
	} `json:"fronts,omitempty"`
	// Event name
	EventName string `json:"event_name,omitempty"`
	// Event ID
	EventId string `json:"event_id,omitempty"`
	// Finishing time
	End string `json:"end,omitempty"`
} 
