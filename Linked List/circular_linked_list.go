package main

import "fmt"

func main(){
	cll := new(CLL)
	
	fmt.Println("Inserting Beginning")
	cll.InsertBeginning(5)
	cll.InsertBeginning(4)
	cll.InsertBeginning(3)
	cll.InsertBeginning(2)
	cll.Display()

	fmt.Println("Inserting End")
	cll.InsertEnd(7)
	cll.InsertEnd(8)
	cll.InsertEnd(9)
	cll.InsertEnd(10)
	cll.InsertEnd(11)
	cll.InsertEnd(12)
	cll.Display()

	fmt.Println("Insert")
	cll.Insert(1,1)
	cll.Insert(6,6)
	cll.Insert(13,13)
	cll.Display()

	fmt.Println("Delete First")
	cll.DeleteFirst()
	cll.Display()

	fmt.Println("Delete End")
	cll.DeleteEnd()
	cll.Display()

	fmt.Println("Testing with new List")
	newCLL := new(CLL)
	newCLL.InsertBeginning(1)
	newCLL.Insert(2,2)
	newCLL.InsertEnd(3)
	newCLL.Insert(4,4)
	newCLL.Display()
	newCLL.Delete(3)
	newCLL.Display()
}

type ListNode struct {
	data interface{}
	next *ListNode
}

type CLL struct {
	head *ListNode
	size int
}

func (cll *CLL) Count() int {
	current := cll.head
	count := 1
	if cll.head == nil {
		return 0
	}
	current = current.next
	for current != cll.head {
		current = current.next
		count++
	}
	return count
}

func (cll *CLL) Display(){
	current := cll.head
	for i := 0; i < cll.size; i++ {
		fmt.Printf("%+v", current.data)
		fmt.Printf(" -> ")
		current = current.next
	}
	fmt.Println()
}

func (cll *CLL) CheckEmptyAndAdd(newNode *ListNode) bool {
	if cll.head != nil {
		return false
	}
	cll.head = newNode
	cll.size++
	return true
} 

func (cll *CLL) InsertBeginning(data interface{}){
	newNode := &ListNode{
		data: data,
		next: nil,
	}
	newNode.next = newNode

	if !cll.CheckEmptyAndAdd(newNode) {
		current := cll.head
		for current.next != cll.head {
			current = current.next
		}

		current.next = newNode
		newNode.next = cll.head
		cll.head = newNode
		cll.size++
	}
}

func (cll *CLL) InsertEnd(data interface{}) {
	newNode := &ListNode{
		data: data,
		next: nil,
	}
	newNode.next = newNode

	if !cll.CheckEmptyAndAdd(newNode){
		
		current := cll.head
		for current.next != cll.head {
			current = current.next
		}
		
		current.next = newNode
		newNode.next = cll.head
		cll.size++
	}
}

func (cll *CLL) Insert(data interface{}, pos int) {
	if pos < 0 || pos > cll.size + 1 {
		fmt.Println("Invalid Position")
		return
	}

	newNode := &ListNode{
		data: data,
		next: nil,
	}
	newNode.next = newNode

	if !cll.CheckEmptyAndAdd(newNode) {
		
		if pos == 1 {
			cll.InsertBeginning(data)
			return
		}

		current := cll.head

		for pos-1 != 1 {
			pos--
			current = current.next
		}

		newNode.next = current.next
		current.next = newNode

		cll.size++
	}
}

func (cll *CLL) DeleteFirst(){
	
	if cll.head == nil {
		fmt.Println("Empty List")
		return
	}

	current := cll.head
	for current.next != cll.head {
		current = current.next
	}

	if current == cll.head {
		cll.head = nil
	} else {
		current.next = cll.head.next
		cll.head = current.next
	}
	cll.size--
}


func (cll *CLL) DeleteEnd(){
	if cll.head == nil {
		fmt.Println("Empty List")
		return
	}

	current := cll.head
	var prev *ListNode = nil
	for current.next != cll.head {
		prev = current
		current = current.next
	}

	if prev == nil {
		cll.head = nil
	} else {
		prev.next = current.next
	}
	cll.size--
}

func (cll *CLL) Delete(pos int) {
	if cll.head == nil {
		fmt.Println("Empty List")
		return
	}

	if pos < 1 || pos > cll.size {
		fmt.Println("Invlaid Position")
		return
	}

	if pos == 1 {
		cll.DeleteFirst()
		return
	}

	current := cll.head
	var prev *ListNode = nil

	for pos != 1 {
		prev = current
		current = current.next
		pos--
	}

	prev.next = current.next
	cll.size--
}























































































