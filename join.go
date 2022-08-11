package piscine

func Join(strs []string, sep string) string {
	count := ""
	sep_s := sep
	for index, i := range strs {
		count += i
		if index != len(strs)-1 {
			count += sep_s
		}
	}
	return count
}
