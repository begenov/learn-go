package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || nb == 1 {
		return 0
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		count := 1
		for i := 0; i < power; i++ {
			count *= nb
		}
		return count
	}
}
