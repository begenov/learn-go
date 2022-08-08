package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	count := nb - 1
	for i := count; i >= 1; i-- {
		nb *= i
		if nb < 0 {
			return 0
		}
	}
	return nb
}
