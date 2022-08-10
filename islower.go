package piscine

func IsLower(s string) bool {
	count := 0
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			count++
		}
	}
	return count == len(s)
}
