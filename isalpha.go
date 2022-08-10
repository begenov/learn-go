package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if (i < '0' || i > '9') && (i < 'A' || i > 'Z') && (i < 'a' || i > 'z') {
			return false
		}
	}
	return true
}
