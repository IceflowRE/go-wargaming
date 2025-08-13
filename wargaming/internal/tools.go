package internal

import (
	"strconv"
	"strings"
)

func SliceIntToString(slice []int, sep string) string {
	text := make([]string, len(slice))
	for idx, val := range slice {
		text[idx] = strconv.Itoa(val)
	}

	return strings.Join(text, sep)
}
