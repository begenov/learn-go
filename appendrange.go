package piscine

func AppendRange(min, max int) []int {
	var num1 []int
	for num := min; num < max; num++ {
		num1 = append(num1, num)
	}
	return num1
}
