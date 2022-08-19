package piscine

func Unmatch(a []int) int {
	for _, value := range a {
		count := 0
		for _, el := range a {
			if value == el {
				count++
			}
		}
		if count%2 != 0 {
			return value
		}
	}
	return -1
}
