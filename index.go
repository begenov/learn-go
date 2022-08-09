package main

import (
	"fmt"
)
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 0
	}
	a := 0
	for i := range s {
		if s[i] == toFind[a] {
			if a == len(toFind)-1 {
				return i - (len(toFind) - 1)
			}
			a++
		}
	}
	return 0
}
