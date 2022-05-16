package wgn

type WgtvTags struct {
	// List of projects
	Projects struct {
		// Game project ID
		ProjectId int `json:"project_id,omitempty"`
		// Sort order
		Order int `json:"order,omitempty"`
		// Name of the game project
		Name string `json:"name,omitempty"`
	} `json:"projects,omitempty"`
	// List of programs
	Programs struct {
		// Program ID
		ProgramId int `json:"program_id,omitempty"`
		// Sort order
		Order int `json:"order,omitempty"`
		// Localized name of the program
		Name string `json:"name,omitempty"`
	} `json:"programs,omitempty"`
	// List of categories
	Categories struct {
		// Sort order
		Order int `json:"order,omitempty"`
		// Localized name of the category
		Name string `json:"name,omitempty"`
		// Content category ID
		CategoryId int `json:"category_id,omitempty"`
	} `json:"categories,omitempty"`
} 
