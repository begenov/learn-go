package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(Atoi("12345"))
// 	fmt.Println(Atoi("0000000012345"))
// 	fmt.Println(Atoi("012 345"))
// 	fmt.Println(Atoi("Hello World!"))
// 	fmt.Println(Atoi("+1234"))
// 	fmt.Println(Atoi("-1234"))
// 	fmt.Println(Atoi("++1234"))
// 	fmt.Println(Atoi("--1234"))
// }

// func Atoi(s string) int {
// 	count := 0
// 	sign := 1
// 	for i := range s {
// 		if s[i] < '0' || s[i] > '9' {
// 			return 0
// 		}
// 	}
// 	n := []rune(s)
// 	for i := 0; i < len(s); i++ {
// 		if s[i] > '0' || s[i] < '9' {
// 			count *= 10
// 			count += int(n[i] - 48)
// 		}

// 	}
// 	return count
// }

func Atoi(s string) int {
	count := 0
	// for k := 0; k < len(s)-1; k++ {
	// 	if s[k] == '+' || s[k] == '-' {
	// 		continue
	// 	}
	// }

	for i := range s {
		if (s[0] == '-' || s[0] == '+') && i == 0 {
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		count *= 10
		count += int(s[i] - 48)
	}
	if s[0] == '-' {
		count *= -1
	}
	return count
}
