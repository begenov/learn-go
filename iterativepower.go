package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || nb == 1 {
		return 1
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		res := 1
		for i := 0; i < power; i++ {
			res *= nb
		}
		return res
	}
}
