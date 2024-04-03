package list

import "errors"

type Circular struct {
	Head *Node
}

func NewCircular() *Circular {
	return &Circular{}
}

func (c *Circular) Enqueue(value int) {
	node := newNode(value)
	if c.Head == nil {
		c.Head = node
		node.Next = node
	} else {
		node.Next = c.Head
		prev := c.Head
		for prev.Next != c.Head {
			prev = prev.Next
		}
		prev.Next = node
	}
}

func (c *Circular) Dequeue() (int, error) {
	if c.Head == nil {
		return 0, errors.New("empty list")
	}
	value := c.Head.Value
	if c.Head.Next == c.Head {
		c.Head = nil
	} else {
		prev := c.Head
		for prev.Next != c.Head {
			prev = prev.Next
		}
		prev.Next = c.Head.Next
		c.Head = c.Head.Next
	}
	return value, nil
}

func (c *Circular) Peek() (int, error) {
	if c.Head == nil {
		return 0, errors.New("empty list")
	}
	return c.Head.Value, nil
}
