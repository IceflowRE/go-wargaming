package wotb

type EncyclopediaAchievements struct {
	// Section
	Section string `json:"section,omitempty"`
	// Order of achievements
	Order int `json:"order,omitempty"`
	// Service Record
	Options struct {
		// Achievement name
		Name string `json:"name,omitempty"`
		// Link to large image
		ImageBig string `json:"image_big,omitempty"`
		// Image link
		Image string `json:"image,omitempty"`
		// Achievement description
		Description string `json:"description,omitempty"`
	} `json:"options,omitempty"`
	// Achievement name
	Name string `json:"name,omitempty"`
	// Link to large image
	ImageBig string `json:"image_big,omitempty"`
	// Image link
	Image string `json:"image,omitempty"`
	// Achievement description
	Description string `json:"description,omitempty"`
	// Condition
	Condition string `json:"condition,omitempty"`
	// Achievement ID
	AchievementId string `json:"achievement_id,omitempty"`
} 
