package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	n1 := Atoi(args[0])
	n2 := Atoi(args[2])
	ar := args[1]
	switch ar {
	case "+":
		str := n1 + n2
		res := Itoa(str)
		print(res)
	case "-":
		str := n1 - n2
		res := Itoa(str)
		print(res)
	case "*":
		str := n1 * n2
		res := Itoa(str)
		print(res)
	case "/":
		str := n1 / n2
		res := Itoa(str)
		print(res)
	case "%":
		str := n1 % n2
		res := Itoa(str)
		print(res)
	}
}

func print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func Atoi(s string) int {
	flag := false
	if s[0] == '-' {
		flag = true
		s = s[1:]
	}
	res := 0
	for i := 0; i < len(s); i++ {
		// if s[i] == '-' {
		// 	continue
		// }
		res = res*10 + int(rune(s[i]-48))
	}
	if flag {
		res *= -1
	}
	return res
}

func Itoa(n int) string {
	flag := false
	if n < 0 {
		flag = true
		n = n * -1
	}
	res := ""
	for n > 0 {
		res = string(rune(n%10)+48) + res
		n = n / 10
	}
	if flag {
		res = "-" + res
	}
	return res
}
