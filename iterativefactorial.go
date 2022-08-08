package piscine

func IterativeFactorial(nb int) int {
	count := 1
	for i := 0; i < nb-1; i++ {
		count *= nb
	}
	return count
}
