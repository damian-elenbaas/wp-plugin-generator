package generator

import "unicode"

func splitOnCapitals(s string) []string {
	var parts []string
	last := 0

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			parts = append(parts, s[last:i])
			last = i
		}
	}
	parts = append(parts, s[last:])
	return parts
}
