package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	count := 0
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
	}
	n := []rune(s)
	for i := 0; i < len(s); i++ {

		count *= 10
		count += int(n[i] - 48)
	}
	return count
}
