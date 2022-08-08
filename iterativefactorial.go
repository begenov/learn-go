package piscine

func IterativeFactorial(nb int) int {
	count := 0

	for i := 1; i < nb+1; i++ {
		count *= i
	}
	return count
}
