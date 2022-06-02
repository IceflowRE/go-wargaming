package internal

import (
	"strconv"
	"strings"
)

func MapValuesToSlice[K comparable, T any](dict map[K]T) []T {
	values := make([]T, 0, len(dict))
	for _, val := range dict {
		values = append(values, val)
	}
	return values
}

func SliceIntToString(slice []int, sep string) string {
	text := make([]string, len(slice))
	for i, v := range slice {
		text[i] = strconv.Itoa(v)
	}

	return strings.Join(text, sep)
}
