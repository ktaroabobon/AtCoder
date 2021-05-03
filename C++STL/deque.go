package main

import "fmt"

func main() {
	// Create a new Deque
	deque := NewDeque()

	// Inject two items in back
	deque.Append("Second")
	deque.Append("Third")

	// Push an item in front
	deque.AppendLeft("First")

	fmt.Println(deque.Items)

	// Remove an item in front
	deque.PopLeft()

	// Remove an item in back
	deque.Pop()

	// Check if the deque is empty and his values
	fmt.Println(deque.IsEmpty(), deque.Items)
}

func NewDeque() *Deque {
	return &Deque{}
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) AppendLeft(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Append(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) PopLeft() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Pop() interface{} {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}
