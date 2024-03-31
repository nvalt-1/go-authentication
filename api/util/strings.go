package util

import "strings"

func EndsWith(str string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(str, s) {
			return true
		}
	}
	return false
}
