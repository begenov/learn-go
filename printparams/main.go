package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for _, i := range args[1:] {
		for _, j := range i {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
