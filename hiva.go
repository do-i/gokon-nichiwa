package main

import (
	"fmt"
	)

// Ch 9 :- Nine or Hiva in Tongan Language

// Structs and Interfaces
type Node struct {
	value int32
	next *Node
}

func shortList() {
	// root is pointer
	root := new(Node)
	root.value = 69

	child := new(Node)
	root.next = child
	child.value = 13

	// another way to create Node object
	// then assign address of grandChild object to the pointer
	grandChild := Node{888, nil}
	child.next = &grandChild

	printer(root)
}

func printer(list *Node) {
	current := list
	for current.next != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println(current.value)
}

func main() {
	mainHello()
	mainNihao()
	mainHachi()
	fmt.Println("empty")
	shortList()
}
