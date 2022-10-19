// Auto generated file!

package wgn

type WgtvVehiclesOptions struct {
	// Content category ID. Maximum limit: 100.
	CategoryId []int `json:"category_id,omitempty"`
	// Program ID. Maximum limit: 100.
	ProgramId []int `json:"program_id,omitempty"`
	// Game project ID. Maximum limit: 100.
	ProjectId []int `json:"project_id,omitempty"`
}

type WgtvVehicles struct {
	// Vehicle IDs by games
	Vehicles map[string]string `json:"vehicles,omitempty"`
}
