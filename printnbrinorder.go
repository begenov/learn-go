package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	a := []int{}
	if n == 0 {
		a = append(a, 0)
	}
	for n > 0 {
		b := n % 10
		a = append(a, b)
		n = n / 10
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for i := 0; i < len(a); i++ {
		z01.PrintRune(rune(a[i]) + 48)
	}
}
