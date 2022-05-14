package utils

import (
	"strconv"
	"strings"
)

func Contains[T comparable](elems []T, v T) bool {
	for _, e := range elems {
		if v == e {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](elems []T, values []T) bool {
	for _, v := range values {
		if !Contains(elems, v) {
			return false
		}
	}
	return true
}

func SliceIntToString(slice []int, sep string) string {
	text := make([]string, len(slice))
	for i, v := range slice {
		text[i] = strconv.Itoa(v)
	}

	return strings.Join(text, sep)
}
