package list

type Stack struct {
	LinkedList *LinkedList
}

func NewStack() *Stack {
	return &Stack{LinkedList: NewLinkedList()}
}

func (s *Stack) Push(value int) {
	s.LinkedList.InsEnd(value)
}

func (s *Stack) Pop() (int, error) {
	return s.LinkedList.DelEnd()
}

func (s *Stack) Peek() (int, error) {
	return s.LinkedList.PeekEnd()
}

func (s *Stack) IsEmpty() bool {
	return s.LinkedList.IsEmpty()
}

func (s *Stack) Length() int {
	return s.LinkedList.Length
}
