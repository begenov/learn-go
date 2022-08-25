package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func validateOperator(test string) bool {
	op := []string{"+", "-", "*", "/", "%"}
	for _, res := range op {
		if res == test {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		fmt.Print()
	} else {
		if !validateOperator(args[1]) {
			fmt.Println(0)
		} else {
			premier, _ := strconv.Atoi(args[0])
			second, _ := strconv.Atoi(args[2])

			if args[1] == "%" && second == 0 {
				fmt.Print("No Modulo by 0\n")
			} else if args[1] == "/" && second == 0 {
				printstr()
			} else if args[1] == "+" {
				fmt.Println(premier + second)
			} else if args[1] == "-" {
				fmt.Println(premier - second)
			} else if args[1] == "*" {
				fmt.Println(premier * second)
			} else if args[1] == "/" {
				fmt.Println(premier / second)
			} else {
				fmt.Println(premier % second)
			}

		}
	}
}

func printstr() {
	a := "No modulo by 0"
	for _, i := range a {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
