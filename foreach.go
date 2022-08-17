package piscine

// import "fmt"

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}

// func oraz(n int) {
// 	count := '0'
// 	if n == 0 {
// 		fmt.Println(count)
// 		return

// 	}
// 	for i := 1; i <= n%10; i++ {
// 		count++
// 	}
// 	for j := -1; j >= n%10; j-- {
// 		count++
// 	}
// 	if n/10 != 0 {
// 		oraz(n / 10)
// 	}
// 	fmt.Println(count)
// }

// func PrintNbr(n int) {
// 	count := 1
// 	for i := 1; i < n; i++ {
// 		count++
// 	}
// 	fmt.Print(count)
// }
