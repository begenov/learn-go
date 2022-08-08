package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || nb == 1 {
		return 0
	} else if 1 == power {
		return 1
	} else {
		res := 1
		for i := 0; i < power; i++ {
			res *= nb
		}
		return res
	}
}
