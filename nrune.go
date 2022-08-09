package piscine

func NRune(s string, n int) rune {
	rune := []rune(s)
	count := len(s)
	for index := range rune {
		count = index
	}
	if n-1 >= 0 && n-1 <= count {
		return rune[n-1]
	}
	return 0
}
