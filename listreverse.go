package piscine

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, 2)
// 	ListPushBack(link, 3)
// 	ListPushBack(link, 4)

// 	ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	var massivInterface []interface{}

	for i := l.Head; i != nil; i = i.Next {
		massivInterface = append(massivInterface, i.Data)
	}

	links := l.Head

	for i := len(massivInterface) - 1; i >= 0; i-- {
		links.Data = massivInterface[i]
		links = links.Next
	}
}

// func ListPushBack(l *List, data interface{}) {
// 	n := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = n
// 		l.Tail = n
// 	} else {
// 		l.Tail.Next = n
// 		l.Tail = n
// 	}
// }
