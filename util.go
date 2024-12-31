package overlay

import "strings"

// clamp calculates the lowest possible number between the given boundaries.
func clamp(v, lower, upper int) int {
	return min(max(v, lower), upper)
}

// lines normalises any non standard new lines within a string, and then splits and returns a slice
// of strings split on the new lines.
func lines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(s, "\n")
}

// whitescpace returns a string of whitespace characters of the requested length.
func whitespace(length int) string {
	return strings.Repeat(" ", length)
}
