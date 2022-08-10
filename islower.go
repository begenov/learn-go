package piscine

func IsLower(s string) bool {
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			return false
		} else if i <= 'a' || i >= 'z' {
			return false
		}
	}
	return true
}
