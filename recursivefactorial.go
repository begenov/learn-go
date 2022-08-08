package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}

	fac := nb * RecursiveFactorial(nb-1)
	if fac < 0 {
		return 0
	}
	return fac
}
