package main

import (
	"fmt"
)

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	}
	return nb * RecursiveFactorial(nb-1)
}
