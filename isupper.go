package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}

func IsUpper(s string) bool {
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			return false
		} else if i <= 'A' || i >= 'Z' {
			return false
		}
	}
	return true
}
