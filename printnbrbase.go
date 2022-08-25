package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	k := 0
	if nbr == -9223372036854775808 {
		nbr += 1
		k += 1
	}
	if checkBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	st := []rune(base)
	var q []int
	if nbr < 0 {
		nbr *= -1
		z01.PrintRune('-')
	}
	mod := 0
	t := len(base)
	for nbr > 0 {
		mod = nbr % t
		q = append(q, mod)
		nbr = int(nbr / t)
	}

	for i := len(q) - 1; i >= 0; i-- {
		if i == 0 && k == 1 {
			q[i] += 1
		}
		z01.PrintRune(st[q[i]])
	}
}

func checkBase(base string) bool {
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
