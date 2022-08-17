package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	f1 := true
	f2 := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			f1 = false
		} else if f(a[i], a[i+1]) < 0 {
			f2 = false
		}
	}
	return f1 || f2
}
