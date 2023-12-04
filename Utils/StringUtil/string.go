package StringUtil

import "strings"

func ContainsString(slice []string, value string) bool {
	for _, s := range slice {
		if strings.Contains(s, value) {
			return true
		}
	}
	return false
}
