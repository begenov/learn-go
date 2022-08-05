package piscine

func StrRev(s string) string {
	oa := []rune(s)
	for i, q := 0, len(oa)-1; 1 < q; i, q = i+1, q-1 {
		oa[i], oa[q] = oa[q], oa[i]
	}
	return string(oa)
}
