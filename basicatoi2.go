package piscine

func BasicAtoi2(s string) int {
	count := 0
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
	}
	n := []rune(s)
	for i := 0; i < len(s); i++ {

		count *= 10
		count += int(n[i] - 48)
	}
	return count
}
