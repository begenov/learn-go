package main

import (
	"fmt"
)

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}

// func CollatzCountdown(start int) int {
// 	count := 0
// 	if start <= 0 {
// 		return -1
// 	} else if start == 1 {
// 		return count
// 	} else if start%2 == 0 {
// 		count++
// 		return CollatzCountdown(start / 2)
// 	} else if start%2 != 0 {
// 		count++
// 		return CollatzCountdown(start*3 + 1)
// 	}
// 	return count
// }

// func CollatzCountdown(n int) int {
// 	count := 0
// 	if n == 1 {
// 		return count
// 	} else if n <= 0 {
// 		return -1
// 	}
// 	if n != 1 {
// 		count++
// 	}
// 	if n%2 == 0 {
// 		return CollatzCountdown(n/2) + 1
// 	}
// 	if n%2 != 0 {
// 		return CollatzCountdown(3*n+1) + 1
// 	}
// 	return count
// }

func CollatzCountdown(n int) int {
	count := 0
	for {
		if n == 1 {
			return count
			break
		} else if n <= 0 {
			return -1
			break
		} else if n%2 == 0 {
			n = n / 2
			count++
		} else if n%2 != 0 {
			n = 3*n + 1
			count++
		}
	}
	return count
}
