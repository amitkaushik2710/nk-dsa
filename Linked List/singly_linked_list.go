package main

import "fmt"

func main(){

	ll := new(LinkedList)
	
	fmt.Println("Insert Beginning")
	ll.InsertBeginning(4)
	ll.InsertBeginning(3)
	ll.Display()

	fmt.Println("Insert End")
	ll.InsertEnd(6)
	ll.Display()

	fmt.Println("Insert At")
	ll.Insert(1,1)
	ll.Insert(2,2)
	ll.Insert(5,5)
	ll.Insert(7,7)
	ll.Display()
}

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) Display(){
	current := ll.head

	for current != nil {
		fmt.Printf(" %+v ->", current.data)
		current = current.next
	}

	fmt.Println()
	fmt.Println("Size : ",ll.size)
	fmt.Println()
}

func (ll *LinkedList) InsertBeginning(data interface{}){
	newNode := &ListNode{
		data: data,
		next: nil,
	}
	newNode.next = ll.head
	ll.head = newNode
	ll.size++
}

func (ll *LinkedList) InsertEnd(data interface{}){
	newNode := &ListNode{
		data: data,
		next: nil,
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	
	if current == nil {
		ll.head = newNode
	} else {
		current.next = newNode
	}

	ll.size++
}


func (ll *LinkedList) Insert(data interface{}, pos int) {
	if pos < 1 || pos > ll.size + 1 {
		fmt.Println("Invalid postition")
		return
	}

	if pos == 1 {
		ll.InsertBeginning(data)
		return
	}

	current := ll.head
	for pos - 1 > 1 {
		current = current.next
		pos--
	}

	newNode := &ListNode{
		data: data,
		next: current.next,
	}
	current.next = newNode
	ll.size++
}




















