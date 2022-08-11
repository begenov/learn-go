package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 1; i-- {
		for _, j := range args[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
