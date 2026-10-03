package main

import "fmt"

func main(){
	
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println("Pop : ", s.Pop())
	fmt.Println("Pop : ", s.Pop())
	fmt.Println("Pop : ", s.Pop())
	fmt.Println("Pop : ", s.Pop())
	fmt.Println("Pop : ", s.Pop())
}

type ListNode struct {
	data interface{}
	next *ListNode
}

type Stack struct {
	head *ListNode
	size int
}

const (
       STACK_SIZE = 10
)

// Is FULL
func (s *Stack) IsFull() bool {
	return s.size >= STACK_SIZE
}

// Is Empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Push
func (s *Stack) Push(data interface{}){
	if s.IsFull() {
		fmt.Println("Stack is Full")
		return
	}

	newNode := &ListNode{
		data: data,
		next: nil,
	}

	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}

	s.size++
}

// Pop
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty")
		return "Empty"
	}

	data := s.head.data
	s.head = s.head.next
	s.size--

	return data
}

// Size
func (s *Stack) Size() int {
	return s.size
}
