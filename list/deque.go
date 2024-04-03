package list

type Deque struct {
	DoublyLinkedList *DoublyLinkedList
}

func NewDeque() *Deque {
	return &Deque{DoublyLinkedList: NewDoublyLinkedList()}
}

func (d *Deque) PushFront(value int) {
	d.DoublyLinkedList.InsBeg(value)
}

func (d *Deque) PushBack(value int) {
	d.DoublyLinkedList.InsEnd(value)
}

func (d *Deque) PopFront() (int, error) {
	return d.DoublyLinkedList.DelBeg()
}

func (d *Deque) PopBack() (int, error) {
	return d.DoublyLinkedList.DelEnd()
}

func (d *Deque) PeekFront() (int, error) {
	return d.DoublyLinkedList.PeekBeg()
}

func (d *Deque) PeekBack() (int, error) {
	return d.DoublyLinkedList.PeekEnd()
}

func (d *Deque) IsEmpty() bool {
	return d.DoublyLinkedList.IsEmpty()
}

func (d *Deque) Length() int {
	return d.DoublyLinkedList.Length
}
