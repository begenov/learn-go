package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	var validate bool
	num := 0

	if s[0] == '-' {
		s = s[1:]
		validate = true
	}
	for _, w := range s {

		if w < '0' || w > '9' {
			os.Exit(0)
		}

		num = num*10 + int(w) - '0'
	}
	if validate {
		num = num * -1
	}
	return num
}

func IsNumeric(s string) bool {
	for _, letter := range s {
		if letter < '0' || letter > '9' {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}

	num1 := Atoi(arg[0])
	num2 := Atoi(arg[2])
	op := arg[1]

	if (num1 > 20000000 || num2 > 20000000) || (num1 < -20000000 || num2 < -20000000) {
		return
	}

	if num1 > 0 && num2 > 0 && num1+num2 < 0 {
		return
	}

	if IsNumeric(arg[0]) || IsNumeric(arg[2]) {
		switch op {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			if num2 == 0 {
				fmt.Println("No division by 0")
				return
			}
			fmt.Println(num1 / num2)
		case "%":
			if num2 == 0 {
				fmt.Println("No modulo by 0")
				return
			}
			fmt.Println(num1 % num2)
		}
	} else {
		return
	}
}
