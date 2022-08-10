package piscine

func IsPrintable(s string) bool {
	arr := []rune(s)
	for i := range arr {
		if arr[i] >= 0 && arr[i] <= 31 {
			return false
		}
	}
	return true
}
