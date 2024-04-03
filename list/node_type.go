package list

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}
