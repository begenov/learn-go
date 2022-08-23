package piscine

// package main

// import (
// 	"fmt"
// )

// func PrintElem(node *NodeL) {
// 	fmt.Println(node.Data)
// }

// func StringToInt(node *NodeL) {
// 	node.Data = 2
// }

// func PrintList(l *List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, "->")
// 		it = it.Next
// 	}
// 	fmt.Print("nil", "\n")
// }

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, "hello")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "there")
// 	ListPushBack(link, 23)
// 	ListPushBack(link, "!")
// 	ListPushBack(link, 54)

// 	PrintList(link)

// 	fmt.Println("--------function applied--------")
// 	ListForEachIf(link, PrintElem, IsPositiveNode)

// 	ListForEachIf(link, StringToInt, IsAlNode)

// 	fmt.Println("--------function applied--------")
// 	PrintList(link)

// 	fmt.Println()
// }

func IsNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return true
	case string, rune:
		return false
	}
	return false
}

func ListForEachIf(l *List, f func(*NodeL), comp func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if comp(it) {
			f(it)
		}
		it = it.Next
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

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
