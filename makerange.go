package piscine

func MakeRange(min, max int) []int {
	l := max - min
	if max <= min {
		return nil
	}
	as := make([]int, l)
	for i := 0; i < l; i++ {
		as[i] = i + min
	}
	return as
}
