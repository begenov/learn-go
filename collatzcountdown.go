package piscine

func CollatzCountdown(n int) int {
	count := 0
	if n == 1 {
		return count
	} else if n <= 0 {
		return -1
	}
	if n != 1 {
		count++
	}
	if n%2 == 0 {
		return CollatzCountdown(n/2) + 1
	}
	if n%2 != 0 {
		return CollatzCountdown(3*n+1) + 1
	}
	return count
}
