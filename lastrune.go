package piscine

func LastRune(s string) rune {
	a := []rune(s)
	b := 0
	for i := range s {
		b = i + 1
	}
	return a[b-1]
}
