package iteration

import "strings"

// Repeat repeats a given string the supplied number of times.
// For example, Repeat("abc", 3) returns "abcabcabc".
// The function doesn't copy, but uses `(strings.Builder).WriteString` to achieve the required functionality.
func Repeat(s string, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteString(s)
	}

	return repeated.String()
}
