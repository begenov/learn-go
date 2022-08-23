package piscine

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "man")
// 	ListPushFront(link, "how are you")

// 	it := link.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " ")
// 		it = it.Next
// 	}
// 	fmt.Println()
// }

// func ListPushFront(l *List, data interface{}) {
// 	nq := new(NodeL)
// 	n := &NodeL{Data: data}
// 	if n == nil {
// 		l.Head = n
// 		l.Tail = n
// 	} else {
// 		l.Head = n
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
// package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL{Data: data}, l.Head
	} else {
		newNode := &NodeL{Data: data}
		newNode.Next, l.Head = l.Head, newNode
	}
}
