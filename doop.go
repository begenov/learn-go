package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	a := strtoint(args[0])
	b := strtoint(args[2])
	// fmt.Println(a, b)
	arr := args[1]
	switch arr {
	case "+":
		str := a + b
		fmt.Println(str)

	case "-":
		str := a - b
		fmt.Println(str)
	case "*":
		str := a * b
		fmt.Println(str)

	case "/":
		str := a / b
		fmt.Println(str)

	default:
		fmt.Println("as")
	}
}

func print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

func strtoint(s string) int {
	count := 0
	for _, v := range s {
		if v == '-' {
			continue
		} else {
			count = count*10 + int(v-48)
		}
	}
	// fmt.Println(count)
	if s[0] == '-' {
		count *= -1
	}
	// fmt.Println(count)

	return count
}

func inttostr(num int) string {
	res := ""
	if num == 0 {
		return "0"
	}
	for num != 0 {
		res += string(rune(num%10 + 48))
		num /= 10
	}
	as := []rune(res)
	arr := ""
	for i := len(as) - 1; i >= 0; i-- {
		arr += string(as[i])
	}
	// fmt.Println(res)
	return arr
}
