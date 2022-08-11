package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		for j := 0; j < len(arg)-1; j++ {
			if arg[j] > arg[j+1] {
				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	for _, v := range arg {
		for _, c := range v {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
