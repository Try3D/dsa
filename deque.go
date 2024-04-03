package main

import "errors"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type Deque struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) PushFront(value int) {
	node := NewNode(value)
	if d.Head == nil {
		d.Head = node
		d.Tail = node
	} else {
		node.Next = d.Head
		d.Head.Prev = node
		d.Head = node
	}
	d.Length++
}

func (d *Deque) PushBack(value int) {
	node := NewNode(value)
	if d.Tail == nil {
		d.Head = node
		d.Tail = node
	} else {
		node.Prev = d.Tail
		d.Tail.Next = node
		d.Tail = node
	}
	d.Length++
}

func (d *Deque) PopFront() (int, error) {
	if d.Head == nil {
		return 0, errors.New("Deque is empty")
	}
	d.Length--
	value := d.Head.Value
	d.Head = d.Head.Next
	if d.Head != nil {
		d.Head.Prev = nil
	} else {
		d.Tail = nil
	}
	return value, nil
}

func (d *Deque) PopBack() (int, error) {
	if d.Tail == nil {
		return 0, errors.New("Deque is empty")
	}
	d.Length--
	value := d.Tail.Value
	d.Tail = d.Tail.Prev
	if d.Tail != nil {
		d.Tail.Next = nil
	} else {
		d.Head = nil
	}
	return value, nil
}

func (d *Deque) Front() (int, error) {
	if d.Head == nil {
		return 0, errors.New("Deque is empty")
	}
	return d.Head.Value, nil
}

func (d *Deque) Back() (int, error) {
	if d.Tail == nil {
		return 0, errors.New("Deque is empty")
	}
	return d.Tail.Value, nil
}

func (d *Deque) Rotate(n int) {
	if n == 0 {
		return
	}
	if n > 0 {
		for i := 0; i < n; i++ {
			d.PushBack(d.Head.Value)
			d.PopFront()
		}
	} else {
		for i := 0; i < -n; i++ {
			d.PushFront(d.Tail.Value)
			d.PopBack()
		}
	}
}

func (d *Deque) Len() int {
	return d.Length
}

func (d *Deque) Clear() {
	d.Head = nil
	d.Tail = nil
	d.Length = 0
}

func (d *Deque) Index(n int) (int, error) {
	if n >= d.Length || n < 0 {
		return 0, errors.New("Index out of range")
	}
	node := d.Head
	for i := 0; i < n; i++ {
		node = node.Next
	}
	return node.Value, nil
}

func (d *Deque) Insert(n, value int) error {
	if n >= d.Length || n < 0 {
		return errors.New("Index out of range")
	}
	node := d.Head
	for i := 0; i < n; i++ {
		node = node.Next
	}
	newNode := NewNode(value)
	newNode.Prev = node.Prev
	newNode.Next = node
	node.Prev.Next = newNode
	node.Prev = newNode
	d.Length++
	return nil
}

func (d *Deque) Search(value int) (int, error) {
	node := d.Head
	for i := 0; i < d.Length; i++ {
		if node.Value == value {
			return i, nil
		}
		node = node.Next
	}
	return 0, errors.New("Value not found")
}
