package list

import (
	"errors"
)

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) InsBeg(value int) {
	node := newNode(value)
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node
		dll.Head = node
	}
	dll.Length++
}

func (dll *DoublyLinkedList) InsEnd(value int) {
	node := newNode(value)
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
	} else {
		dll.Tail.Next = node
		node.Prev = dll.Tail
		dll.Tail = node
	}
	dll.Length++
}

func (dll *DoublyLinkedList) InsPos(value int, index int) error {
	if index < 0 || index > dll.Length {
		return errors.New("index out of bounds")
	} else if index == 0 {
		dll.InsBeg(value)
		return nil
	} else if index == dll.Length {
		dll.InsEnd(value)
		return nil
	} else {
		node := newNode(value)
		prev := dll.Head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		node.Next = prev.Next
		node.Prev = prev
		prev.Next.Prev = node
		prev.Next = node
		dll.Length++
		return nil
	}
}

func (dll *DoublyLinkedList) DelBeg() (int, error) {
	if dll.Head == nil {
		return 0, errors.New("empty list")
	}
	value := dll.Head.Value
	dll.Head = dll.Head.Next
	if dll.Head == nil {
		dll.Tail = nil
	} else {
		dll.Head.Prev = nil
	}
	dll.Length--
	return value, nil
}

func (dll *DoublyLinkedList) DelEnd() (int, error) {
	if dll.Head == nil {
		return 0, errors.New("empty list")
	}
	value := dll.Tail.Value
	dll.Tail = dll.Tail.Prev
	if dll.Tail == nil {
		dll.Head = nil
	} else {
		dll.Tail.Next = nil
	}
	dll.Length--
	return value, nil
}

func (dll *DoublyLinkedList) DelPos(index int) (int, error) {
	if index < 0 || index >= dll.Length {
		return 0, errors.New("index out of bounds")
	} else if index == 0 {
		return dll.DelBeg()
	} else if index == dll.Length-1 {
		return dll.DelEnd()
	} else {
		prev := dll.Head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		value := prev.Next.Value
		prev.Next = prev.Next.Next
		prev.Next.Prev = prev
		dll.Length--
		return value, nil
	}
}

func (dll *DoublyLinkedList) Search(value int) bool {
	node := dll.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.Head == nil
}

func (dll *DoublyLinkedList) PeekBeg() (int, error) {
	if dll.Head == nil {
		return 0, errors.New("empty list")
	}
	return dll.Head.Value, nil
}

func (dll *DoublyLinkedList) PeekEnd() (int, error) {
	if dll.Head == nil {
		return 0, errors.New("empty list")
	}
	return dll.Tail.Value, nil
}
