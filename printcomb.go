package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := '1'; j <= '8'; j++ {
			for x := '2'; x <= '9'; x++ {
				if i == '7' && j == '8' && x == '9' {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(x))
					z01.PrintRune('\n')
				} else if i < j {
					if j < x {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(x))
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
}
