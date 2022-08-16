package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}
	setPoint(&points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	atoi(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	atoi(points.y)
	z01.PrintRune('\n')
}

func atoi(n int) {
	if n < 10 {
		z01.PrintRune('-')
	} else {
		aas(n)
	}
}

func aas(n int) {
	count := '0'
	if n == 0 {
		z01.PrintRune(rune(count))
		return
	}
	for i := 1; i <= n%10; i++ {
		count++
	}
	for i := -1; i >= n%10; i-- {
		count++
	}
	if n/10 != 0 {
		aas(n / 10)
	}
	z01.PrintRune(rune(count))
}
