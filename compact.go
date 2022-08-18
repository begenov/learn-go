package piscine

func Compact(ptr *[]string) int {
	var arr []string
	count := 0
	for _, i := range *ptr {
		if i != "" {
			arr = append(arr, i)
			count++
		}
	}
	*ptr = arr

	return count
}
