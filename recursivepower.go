package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0s
	} else if power == 0 || power == 1 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

