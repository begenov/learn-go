package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for i := range tab {
		if f(tab[i]) {
			count++
		}
	}
	return count
}
