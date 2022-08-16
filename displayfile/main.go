package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("File name missing")
	} else if len(arg) == 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("Almost there!!")
	}
}
