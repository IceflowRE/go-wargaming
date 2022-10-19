// Auto generated file!

package wows

type EncyclopediaAccountlevelsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use "-" in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string `json:"fields,omitempty"`
}

type EncyclopediaAccountlevels struct {
	// URL to Service Record level icon
	Image *string `json:"image,omitempty"`
	// Points to achieve the Service Record level
	Points *int `json:"points,omitempty"`
	// Service Record level
	Tier *int `json:"tier,omitempty"`
}
