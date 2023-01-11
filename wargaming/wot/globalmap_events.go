// Auto generated file!

package wot

type GlobalmapEventsOptions struct {
	// Event ID
	EventId *string `json:"event_id,omitempty"`
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
	// Language. Default is "en". Valid values:
	//
	// "en" - English (by default)
	// "de" - German
	// "fr" - French
	// "es" - Spanish
	// "pl" - Polish
	// "tr" - Turkish
	Language *string `json:"language,omitempty"`
	// Page limit. Default is 5. Min value is 1. Maximum value: 20.
	Limit *int `json:"limit,omitempty"`
	// Page number. Default is 1. Min value is 1.
	PageNo *int `json:"page_no,omitempty"`
	// Response with events filtered by status. Valid values:
	//
	// "PLANNED" - Upcoming event
	// "ACTIVE" - Current event
	// "FINISHED" - Event is over
	Status *string `json:"status,omitempty"`
}

type GlobalmapEvents struct {
	// Finishing time
	End *string `json:"end,omitempty"`
	// Event ID
	EventId *string `json:"event_id,omitempty"`
	// Event name
	EventName *string `json:"event_name,omitempty"`
	// Fronts. Only event fronts are presented in response.
	Fronts []*struct {
		// Front ID
		FrontId *string `json:"front_id,omitempty"`
		// Front name
		FrontName *string `json:"front_name,omitempty"`
		// Link to Front
		Url *string `json:"url,omitempty"`
	} `json:"fronts,omitempty"`
	// Start time
	Start *string `json:"start,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
}
