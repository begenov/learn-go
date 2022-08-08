package piscine

func IterativeFactorial(nb int) int {
	count := 1

	for i := 1; i < nb+1; i++ {
		count *= i
	}
	return count
}
