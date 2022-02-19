package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type singlyLinkedList struct {
	len  int
	head *Node
}

func initList() *singlyLinkedList {

	return &singlyLinkedList{}
}

func (s *singlyLinkedList) addFront(data string) {

	node := &Node{
		data: data,
	}

	if s.head == nil {

		s.head = node

	} else {

		s.head = node
		node.next = s.head

	}

	s.len++
	return
}

//Function to convert singly linked list to circular linked list
func (s *singlyLinkedList) ToCircular() {

	current := s.head

	for current.next != nil {

		current = current.next

	}

	current.next = s.head
}

func main() {
	singleList := initList()
	fmt.Printf("AddFront: D\n")
	singleList.addFront("D")
	fmt.Printf("addFront: C\n")
	singleList.addFront("C")
	fmt.Printf("addFront: B\n")
	singleList.addFront("B")
	fmt.Printf("addFront: A\n")
	singleList.addFront("A")

	fmt.Printf("Size: %d\n", singleList.len)
	singleList.ToCircular()
}
