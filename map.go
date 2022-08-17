package piscine

func Map(f func(int) bool, a []int) []bool {
	count := []bool{}
	for _, j := range a {
		count = append(count, f(j))
	}
	return count
}

// func IsPrime(nb int) bool {
// 	if nb < 0 || nb == 0 || nb == 1 {
// 		return false
// 	} else if nb == 2 || nb == 3 {
// 		return true
// 	}
// 	for i := 2; i < nb; i++ {
// 		count := nb % i
// 		if count == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
