package piscine

func StrLen(s string) int {
	count := 0
	asd := []rune(s)
	for i := 0; i < len(asd); i++ {
		count++
	}
	return count
}
