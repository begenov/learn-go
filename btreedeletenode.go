package piscine

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	piscine.BTreeInsertData(root, "5")
// 	node := BTreeSearchItem(root, "4")
// 	fmt.Println("Before delete:")
// 	piscine.BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		} else {
			temp := BTreeMin(root.Right)

			root.Data = temp.Data
			root.Right = BTreeDeleteNode(root.Right, temp)
		}
	}
	return root
}
