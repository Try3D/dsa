package binarytree

import "fmt"

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(value int) {
	insertHandler(bt.Root, value)
}

func insertHandler(node *Node, value int) {
	if node == nil {
		node = NewNode(value)
	} else if value < node.Value {
		insertHandler(node.Left, value)
	} else if value > node.Value {
		insertHandler(node.Right, value)
	}
}

func (bt *BinaryTree) Delete(value int) bool {
	return deleteHandler(bt.Root, value)
}

func deleteHandler(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value < node.Value {
		return deleteHandler(node.Left, value)
	} else if value > node.Value {
		return deleteHandler(node.Right, value)
	} else {
		if node.Left == nil && node.Right == nil {
			node = nil
		} else if node.Left == nil {
			node = node.Right
		} else if node.Right == nil {
			node = node.Left
		} else {
			min := node.Right
			for min.Left != nil {
				min = min.Left
			}
			node.Value = min.Value
			deleteHandler(node.Right, min.Value)
		}
		return true
	}
}

func (bt *BinaryTree) Search(value int) bool {
	return searchHandler(bt.Root, value)
}

func searchHandler(node *Node, value int) bool {
	if node == nil {
		return false
	} else if node.Value == value {
		return true
	} else if value < node.Value {
		return searchHandler(node.Left, value)
	} else {
		return searchHandler(node.Right, value)
	}
}

func (bt *BinaryTree) PreOrder() {
	preOrderHandler(bt.Root)
}

func preOrderHandler(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Value)
	preOrderHandler(node.Left)
	preOrderHandler(node.Right)
}

func (bt *BinaryTree) InOrder() {
	inOrderHandler(bt.Root)
}

func inOrderHandler(node *Node) {
	if node == nil {
		return
	}
	inOrderHandler(node.Left)
	fmt.Printf("%v ", node.Value)
	inOrderHandler(node.Right)
}

func (bt *BinaryTree) PostOrder() {
	postOrderHandler(bt.Root)
}

func postOrderHandler(node *Node) {
	if node == nil {
		return
	}
	postOrderHandler(node.Left)
	postOrderHandler(node.Right)
	fmt.Printf("%v ", node.Value)
}
