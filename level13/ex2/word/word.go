// Package word provides a couple of functions for working with words
package word

import "strings"

// UseCount calculates word usage in provided string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts words in provided string
func Count(s string) int {
	words := strings.Split(s, " ")

	return len(words)
}
