package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}

func ListPushFront(l *List, data interface{}) {
	nq := new(NodeL)
	n := &NodeL{Data: data}
	if n == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Head = n
		l.Tail = n
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
