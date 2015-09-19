package main

import (
	"fmt"
	)

// Ch 8 Pointers
func whatsDaPointMain() {
	// declare int variable and initialize it
	coolPtr := 8

	// send address(&) of the variable
	whatsDaPoint(&coolPtr)

	// print the variable which was updated by the previous function call
	fmt.Println(coolPtr)
}
// accepts pointer int variable type
func whatsDaPoint(coolPtr *int) {
	// increaments the variable by one to which the pointer points
	*coolPtr++
}

// use new keyword to create pointer to int var
func anotherPoint() {
	// declare pointer int variable (default value zero)
	anotherPtr := new(int)

	// assign 10 to the location pointed by the pointer
	*anotherPtr += 10

	// then pass the pointer
	whatsDaPoint(anotherPtr)

	// print the variable that pointed by the pointer use (*)
	fmt.Println(*anotherPtr)
}

// swap two integers using pointer function
func swapIntegers(x, y *int) {
	t := *x
	*x = *y
	*y = t
}
func swapIntegersMain() {
	x := 4
	y := 8
	fmt.Printf("Before swap : x=%d, y=%d\n", x, y)
	swapIntegers(&x, &y)
	fmt.Printf("After swap  : x=%d, y=%d\n", x, y)
}

func main() {
	whatsDaPointMain()
	anotherPoint()
	swapIntegersMain()
}


