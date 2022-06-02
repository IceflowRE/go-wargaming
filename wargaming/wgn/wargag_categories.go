package wgn


// WargagCategoriesOptions options.
type WargagCategoriesOptions struct {
	// Content category ID
	CategoryId *int
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
}

type WargagCategories struct {
	// Content category ID
	CategoryId *int `json:"category_id,omitempty"`
	// Content category name
	Name *string `json:"name,omitempty"`
	// Content type
	Type_ *string `json:"type,omitempty"`
}
