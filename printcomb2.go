package piscine

import "github.com/01-edu/z01"

// package main

// import (
// 	"github.com/01-edu/z01"
// )

// func main() {
// 	PrintComb2()
// }

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			x := 48 + (i / 10)
			y := 48 + (i % 10)
			z := 48 + (j / 10)
			d := 48 + (j % 10)
			if i <= j {
				z01.PrintRune(rune(x))
				z01.PrintRune(rune(y))
				z01.PrintRune(rune(' '))

				z01.PrintRune(rune(z))
				z01.PrintRune(rune(d))

			}
			if i != 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}
