package util

import "regexp"

func Filter[T ~[]E, E any](s T, f func(E) bool) T {
	// Allocate the same size as the input slice
	r := make(T, 0, len(s))
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(s) {
		return true
	}
	return false
}
