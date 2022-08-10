package piscine

func BasicJoin(elems []string) string {
	count := ""
	for _, i := range elems {
		count += i
	}
	return count
}
