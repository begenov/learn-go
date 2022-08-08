package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	}

	fac := nb * RecursiveFactorial(nb-1)
	if fac < 0 {
		return 0
	} else {
		return fac
	}
}
