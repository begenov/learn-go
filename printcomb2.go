package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for num := 0; num <= 98; num++ {
		for iw := num + 1; iw <= 99; iw++ {
			a := 48 + (num / 10)
			b := 48 + (num % 10)
			c := 48 + (iw / 10)
			d := 48 + (iw % 10)
			if num <= iw {
				z01.PrintRune(rune(a))
				z01.PrintRune(rune(b))
				z01.PrintRune(rune(' '))
				z01.PrintRune(rune(c))
				z01.PrintRune(rune(d))
			}
			if num != 98 {
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			} else {
				z01.PrintRune(rune('\n'))
			}
		}
	}
}
