package piscine

func TrimAtoi(s string) int {
	total := 0
	sign := 1
	strs := []rune(s)
	b := []rune{}
	for _, elem := range strs {
		if elem >= '0' && elem <= '9' {
			b = append(b, elem)
		}
		if elem == '-' && len(b) == 0 {
			sign *= -1
		}

	}
	for _, value := range b {
		d := int(value - 48)
		total += d
		total *= 10
	}
	return (total * sign) / 10
}
