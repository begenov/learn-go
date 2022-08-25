package piscine

func AtoiBase(s string, base string) int {
	if checkBase1(base) {
		return 0
	}
	bassest := []rune(base)
	sum := 0
	st := []rune(s)
	for i := 0; i < len(s); i++ {
		sum += index(st[i], bassest) * pow(len(base), len(s)-1-i)
	}
	return sum
}

func checkBase1(base string) bool {
	b := []rune(base)
	if len(b) < 2 {
		return true
	}
	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] == b[j] {
				return true
			}
			if b[i] == '+' || b[i] == '-' {
				return true
			}
		}
	}
	return false
}

func index(r rune, base []rune) int {
	for i := 0; i < len(base); i++ {
		if base[i] == r {
			return i
		}
	}
	return 0
}

func pow(x, n int) int {
	s := 1
	for i := 0; i < n; i++ {
		s *= x
	}
	return s
}
