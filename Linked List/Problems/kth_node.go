package main

import "fmt"

func main(){

	// ll := new(LinkedList)
	
	// fmt.Println("Insert Beginning")
	// ll.InsertBeginning(4)
	// ll.InsertBeginning(3)
	// ll.Display()

	// fmt.Println("Insert End")
	// ll.InsertEnd(6)
	// ll.Display()

	// fmt.Println("Insert At")
	// ll.Insert(1,1)
	// ll.Insert(2,2)
	// ll.Insert(5,5)
	// ll.Insert(7,7)
	// ll.Insert(8,8)
	// ll.Insert(9,9)
	// ll.Insert(10,10)
	// ll.Display()

	// // ll.FindKthNodeBruteForce(7)
	// // ll.FindKthNodeHashMap(7)
	// ll.FindKthNode(7)

	mm := new(LinkedList)
	mm.Insert(1,1)
	mm.Display()
	removeNthFromEnd(mm.head, 1)
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


// Find kth Node from the end of the linked list
func (ll *LinkedList) FindKthNodeBruteForce(k int) {
	n := 0

	current := ll.head
	for current != nil {
		n++
		current = current.next
	}

	fmt.Println("Total Nodes : ", n)

	if n < k - 1 {
		fmt.Println("Insufficient nodes")
		return
	}

	t := n - k 
	dummy := ll.head
	for i := 0; i < t; i++ {
		dummy = dummy.next
	}

	fmt.Printf("%+v", dummy.data)
	fmt.Println()
}

// Find kth Node from the end of the linked list
func (ll *LinkedList) FindKthNodeHashMap(k int) {
	llHashMap := make(map[interface{}]*ListNode)

	current := ll.head
	for i := 1; current != nil; i++ {
		llHashMap[i] = current
		current = current.next
	}

	n := len(llHashMap)
	fmt.Println("Total Nodes : ", n)
	fmt.Println("HashMap : ", llHashMap)

	if n < k {
		fmt.Println("Insufficient nodes")
		return
	}

	dummy := llHashMap[n-k+1]

	fmt.Printf("%+v", dummy.data)
	fmt.Println()
}

// Find kth Node from the end of the linked list
func (ll *LinkedList) FindKthNode(k int) {

	first, second := ll.head, ll.head

	for i := 1; i < k; i++ {
		if second == nil {
			fmt.Println("Insufficient nodes")
			return
		}
		second = second.next
	}

	for second.next != nil {
		first = first.next
		second = second.next
	}

	fmt.Printf("%+v", first.data)
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    first, second := head, head

    for i := 1; i < n; i++ {
        if second == nil {
            return nil
        }
        second = second.next
    }
	fmt.Println("First : ", first.data)
	fmt.Println("Second : ", second.data)

    var prev *ListNode
    for second.next != nil {
        prev = first
        first = first.next
        second = second.next
    }

	if prev == nil {
		return head.next
	}
	
    prev.next = first.next

    return head
}