package main

import (
	"fmt"
	"os"
)

func operations(num1 int, num2 int, oper string) {
	result := 0
	if oper == "+" || oper == "\"+\"" {
		result = num1 + num2
		if num1 == 9223372036854775807 {
			return
		}
		fmt.Println(result)
	}
	if oper == "-" || oper == "\"-\"" {
		result = num1 - num2
		if num1 == -9223372036854775807 {
			return
		}
		fmt.Println(result)
	}
	if oper == "/" || oper == "\"/\"" {
		if num2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = num1 / num2
		fmt.Println(result)
	}
	if oper == "*" || oper == "\"*\"" {
		result = num1 * num2
		if num1 == 9223372036854775807 {
			return
		}
		fmt.Println(result)
	}
	if oper == "%" || oper == "\"%\"" {
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = num1 % num2
		fmt.Println(result)
	}
}

func main() {
	args := os.Args[1:]
	num1 := 0
	num2 := 0
	u := false
	o := false
	if len(args) == 3 {
		for _, arg1 := range args[0] {
			if arg1 == 45 {
				u = true
				continue
			}
			if (arg1 < '0' || arg1 > '9') && arg1 != 45 {
				return
			}
			num1 = (num1*10 + int((arg1 - 48)))
		}
		if u {
			num1 *= -1
		}
		for _, arg2 := range args[2] {
			if arg2 == 45 {
				o = true
				continue
			}
			if (arg2 < '0' || arg2 > '9') && arg2 != 45 {
				return
			}
			num2 = (num2*10 + int((arg2 - 48)))
		}
		if o {
			num2 *= -1
		}
		operations(num1, num2, args[1])

	}
}
