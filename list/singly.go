package list

import (
	"errors"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func newNode(value int) *Node {
	return &Node{Value: value}
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) InsBeg(value int) {
	node := newNode(value)
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}
	ll.Length++
}

func (ll *LinkedList) InsEnd(value int) {
	node := newNode(value)
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}
	ll.Length++
}

func (ll *LinkedList) InsPos(value int, index int) error {
	if index < 0 || index > ll.Length {
		return errors.New("index out of bounds")
	} else if index == 0 {
		ll.InsBeg(value)
		return nil
	} else if index == ll.Length {
		ll.InsEnd(value)
		return nil
	} else {
		node := newNode(value)
		prev := ll.Head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		node.Next = prev.Next
		prev.Next = node
		ll.Length++
		return nil
	}
}

func (ll *LinkedList) DelBeg() (int, error) {
	if ll.Head == nil {
		return 0, errors.New("empty list")
	}
	value := ll.Head.Value
	ll.Head = ll.Head.Next
	ll.Length--
	return value, nil
}

func (ll *LinkedList) DelEnd() (int, error) {
	if ll.Head == nil {
		return 0, errors.New("empty list")
	}
	if ll.Head.Next == nil {
		return ll.DelBeg()
	}
	node := ll.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	value := node.Next.Value
	node.Next = nil
	ll.Tail = node
	ll.Length--
	return value, nil
}

func (ll *LinkedList) DelPos(index int) (int, error) {
	if index < 0 || index >= ll.Length {
		return 0, errors.New("index out of bounds")
	} else if index == 0 {
		return ll.DelBeg()
	} else if index == ll.Length-1 {
		return ll.DelEnd()
	} else {
		prev := ll.Head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		value := prev.Next.Value
		prev.Next = prev.Next.Next
		ll.Length--
		return value, nil
	}
}

func (ll *LinkedList) Search(value int) bool {
	node := ll.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

func (ll *LinkedList) PeekBeg() (int, error) {
	if ll.Head == nil {
		return 0, errors.New("empty list")
	}
	return ll.Head.Value, nil
}

func (ll *LinkedList) PeekEnd() (int, error) {
	if ll.Head == nil {
		return 0, errors.New("empty list")
	}
	node := ll.Head
	for node.Next != nil {
		node = node.Next
	}
	return node.Value, nil
}
