package piscine

// package main

// import (
// 	"fmt"
// )

// func PrintList(l *NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *NodeI, data int) *NodeI {
// 	n := &NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

// func main() {
// 	var link *NodeI
// 	var link2 *NodeI

// 	link = listPushBack(link, 3)
// 	link = listPushBack(link, 5)
// 	link = listPushBack(link, 7)

// 	link2 = listPushBack(link2, -2)
// 	link2 = listPushBack(link2, 9)

// 	PrintList(SortedListMerge(link2, link))
// }

func SortedListMerge(l1 *NodeI, l2 *NodeI) *NodeI {
	l1 = listSort(l1)
	l2 = listSort(l2)

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data <= l2.Data {
		l1.Next = SortedListMerge(l1.Next, l2)
		return l1
	} else {
		l2.Next = SortedListMerge(l1, l2.Next)
		return l2
	}
}

// type NodeI struct {
// 	Data int
// 	Next *NodeI
// }

func listSort(l *NodeI) *NodeI {
	cmpt := 0
	var first *NodeI

	if l == nil || l.Next == nil {
		return l
	}

	for l != nil {
		next := l.Next
		if cmpt == 0 {
			first = l
			cmpt++
		}

		for next != nil {
			if l.Data > next.Data {
				l.Data, next.Data = next.Data, l.Data
			}
			next = next.Next
		}
		l = l.Next
	}
	return first
}
