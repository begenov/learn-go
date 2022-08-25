package piscine

import "github.com/01-edu/z01"

var pos = [8]int{}

func isSafe(col, row int) bool {
	for i := 0; i < col; i++ {
		prevrow := pos[i]
		if prevrow == row {
			return false
		}
		if row == prevrow+(col-i) || row == prevrow-(col-i) {
			return false
		}
	}
	return true
}

func solve(n int) {
	if n == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(pos[i] + 49))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i < 8; i++ {
			if isSafe(n, i) {
				pos[n] = i
				solve(n + 1)
			}
		}
	}
}

func EightQueens() {
	solve(0)
}
