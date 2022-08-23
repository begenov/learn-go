package piscine

// package main

// import (
// 	"fmt"
// )

// func PrintList(l List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func main() {
// 	link := List{}
// 	link2 := List{}

// 	ListPushBack(link, "a")
// 	ListPushBack(link, "b")
// 	ListPushBack(link, "c")
// 	ListPushBack(link, "d")
// 	fmt.Println("-----first List------")
// 	PrintList(link)

// 	ListPushBack(link2, "e")
// 	ListPushBack(link2, "f")
// 	ListPushBack(link2, "g")
// 	ListPushBack(link2, "h")
// 	fmt.Println("-----second List------")
// 	PrintList(link2)

// 	fmt.Println("-----Merged List-----")
// 	ListMerge(link, link2)
// 	PrintList(link)
// }

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil || l1 == nil {
		return
	}

	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Head
		return
	}

	l1.Tail.Next = l2.Head
}
