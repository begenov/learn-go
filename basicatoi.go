package piscine

func BasicAtoi(s string) int {
	count := 0
	n := []rune(s)
	for i := 0; i < len(s); i++ {

		count *= 10
		count += int(n[i] - 48)
	}
	return count
}
