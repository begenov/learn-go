package main

import (
	"fmt"
)

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 12 {
		return 0
	} else if nb == 0 {
		return 1
	}
	count := nb - 1

	for i := count; i >= 1; i-- {
		nb *= i
	}
	return nb
}
