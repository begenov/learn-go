package main

import (
	"fmt"
)

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}

func Concat(str1 string, str2 string) string {
	a := str1 + str2
	return a
}
