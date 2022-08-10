package piscine

func IsUpper(s string) bool {
	a := []rune(s)
	for _, i := range a {
		if i >= 'a' && i <= 'z' {
			if i >= '!' || i <= '+' {
				return false
			}
		}
	}
	return true
}
