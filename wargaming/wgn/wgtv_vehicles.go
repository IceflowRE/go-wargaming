package wgn


// WgtvVehiclesOptions options.
type WgtvVehiclesOptions struct {
	// Content category ID. Maximum limit: 100.
	CategoryId []int
	// Program ID. Maximum limit: 100.
	ProgramId []int
	// Game project ID. Maximum limit: 100.
	ProjectId []int
}

type WgtvVehicles struct {
	// Vehicle IDs by games
	Vehicles map[string]string `json:"vehicles,omitempty"`
}
