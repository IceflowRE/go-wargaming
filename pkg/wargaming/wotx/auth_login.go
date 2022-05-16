package wotx

type AuthLogin struct {
	// URL where user is redirected for authentication.
	// This URL is returned only if parameter nofollow=1 is passed in.
	Location string `json:"location,omitempty"`
} 
