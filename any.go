package piscine

func Any(f func(string) bool, a []string) bool {
	for _, j := range a {
		if f(j) {
			return true
		}
	}
	return false
}

// func IsNumeric(s string) bool {
// 	for _, i := range s {
// 		if i < '0' || i > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }
