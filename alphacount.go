package piscine

func AlphaCount(s string) int {
	a := []rune(s)

	count := 0
	for _, value := range a {
		if value >= 'A' && value <= 'Z' || value >= 'a' && value <= 'z' {
			count++
		}
	}
	return count
}
