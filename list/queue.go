package list

type Queue struct {
	LinkedList *LinkedList
}

func NewQueue() *Queue {
	return &Queue{LinkedList: NewLinkedList()}
}

func (q *Queue) Enqueue(value int) {
	q.LinkedList.InsEnd(value)
}

func (q *Queue) Dequeue() (int, error) {
	return q.LinkedList.DelBeg()
}

func (q *Queue) Peek() (int, error) {
	return q.LinkedList.PeekBeg()
}

func (q *Queue) IsEmpty() bool {
	return q.LinkedList.IsEmpty()
}

func (q *Queue) Length() int {
	return q.LinkedList.Length
}
