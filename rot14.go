package piscine

func Rot14(s string) string {
	var a string
	for _, w := range s {
		if w >= 'a' && w <= 'z' {
			if w+14 > 'z' {
				a += string(w - 12)
			} else {
				a += string(w + 14)
			}
		} else if w >= 'A' && w <= 'Z' {
			if w+14 > 'Z' {
				a += string(w - 12)
			} else {
				a += string(w + 14)
			}
		} else {
			a += string(w)
		}
	}
	return a
}
