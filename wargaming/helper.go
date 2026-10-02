package wargaming

// Bool is a helper function that allocates a new bool value to store val and returns a pointer to it.
// DEPRECATED: Since Go 1.26 use new(val) instead.
func Bool(val bool) *bool {
	return &val
}

// Int is a helper function that allocates a new int value to store val and returns a pointer to it.
// DEPRECATED: Since Go 1.26 use new(val) instead.
func Int(val int) *int {
	return &val
}

// String is a helper function that allocates a new string value to store val and returns a pointer to it.
// DEPRECATED: Since Go 1.26 use new(val) instead.
func String(val string) *string {
	return &val
}
