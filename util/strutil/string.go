package strutil

import "strings"

// Contains checks if a slice contains a string
func Contains(slice []string, item string) bool {
	item = strings.ToLower(item)
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
