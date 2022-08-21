package piscine

import "github.com/01-edu/z01"

func oraz(n int) {
	count := '0'
	if n == 0 {
		z01.PrintRune(count)
		return

	}
	for i := 1; i <= n%10; i++ {
		count++
	}
	for j := -1; j >= n%10; j-- {
		count++
	}
	if n/10 != 0 {
		oraz(n / 10)
	}
	z01.PrintRune(count)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	oraz(n)
}
