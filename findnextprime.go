package piscine

func Prime(nb int) bool {
	if nb < 0 || nb == 0 || nb == 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	}
	for i := 2; i < nb/2+1; i++ {
		count := nb % i
		if count == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for ; ; nb++ {
		if Prime(nb) {
			return nb
		}
	}
}
