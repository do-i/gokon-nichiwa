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
	grandChild := Node{value:888, next:nil}
	// another way to link by calling add method
	child.add(grandChild)

	// one way to print the list, pass to function
	// printer(root)

	// another way to print the list, by calling function
	root.print()
}

func printer(list *Node) {
	current := list
	for current.next != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println(current.value)
}

// Ch 9.2 Methods
func (node *Node) print() {
	fmt.Println("Method to print.")
	printer(node)
}
// pass by object vs pointer; would it be better to pass by pointer? What's memory usage?
// pass by object copys object ??
func (node *Node) add(child Node) {
	node.next = &child
}

// Ch 9.2.2 Embedded Types:- inheritance
type Person struct {
	firstName string
	lastName string
}
type Cyborg struct {
	Person
	model int32
}
func (p *Person) print() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}
// func (cyborg *Cyborg) print() {
// 	// cyborg.print()  // TODO how can I call parent method????
// 	fmt.Println(cyborg.model)
// }

func embeddedTypesMain() {
	person := new(Person)
	person.firstName = "Kumi"
	person.lastName = "Otowa"
	person.print()

	cyborg := new(Cyborg)
	cyborg.firstName = "testujin"
	cyborg.lastName = "28Go"
	cyborg.model = 28
	cyborg.print()
}

// 9.3 Interfaces
type Animal interface {
	Speak() string
}
type Dog struct {
}
func (dog Dog) Speak() string {
	return "Woof!"
}
type Cat struct {
}
func (cat Cat) Speak() string {
	return "Meow!"
}
type Dev struct {

}
func (dev Dev) Speak() string {
	return "Ship it!"
}

func interfacesMain() {
	animals := []Animal{Dog{}, Cat{}, Dev{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func mainHiva() {
	mainHello()
	mainNihao()
	mainHachi()
	fmt.Println("empty")
	shortList()
	embeddedTypesMain()
	interfacesMain()
}
