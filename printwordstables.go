package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, value := range a {
		for _, j := range value {
			z01.PrintRune(rune(j))
		}
		z01.PrintRune('\n')
	}
}
