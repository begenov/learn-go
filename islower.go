package piscine

import "unicode"

func IsLower(s string) bool {
	for _, i := range s {
		if !unicode.IsLower(i) {
			return false
		}
	}
	return true
}
