package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	} else if nb == 1 {
		return 1
	}
	count := 0
	for i := 0; i < nb; i++ {
		count = i * i
		if count > nb {
			return 0
		} else if nb == count {
			return i
		}
	}
	return 0
}
