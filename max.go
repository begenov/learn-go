package piscine

func Max(a []int) int {
	max := 1
	for _, j := range a {
		if j > max {
			max = j
		}
	}
	return max
}
