package wgn

// WargagTagsOptions options.
type WargagTagsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Tag ID
	TagId *int
}

type WargagTags struct {
	// Tag name
	Name *string `json:"name,omitempty"`
	// Tag ID
	TagId *int `json:"tag_id,omitempty"`
}
