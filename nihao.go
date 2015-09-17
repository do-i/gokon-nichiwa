package main

import (
	"fmt"
	"strconv"
	)

func main() {
	bunchOfNum := []float64{5, 7, 10}
	fmt.Printf("%.2f\n", average(bunchOfNum))
	mainHello()
	multiReturnMain()
	nameReturnTypeMain()
	variadicMain()
}
// input parameter within parenthisis. return type followed
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// name the return type ... ref. page 85
func nameReturnTypeMain() {
	fmt.Println(nameReturnType())
}
func nameReturnType() (result string) {
	// variable result is already declared in the method signature 
	// so you cannot do result := "yoyo!" so just assign a value.
	result = "done!"
	return
	// compiler allows you to return result explicitly, but not necessary
}

// 7.2 return multiple types ... ref. page 86
func multiReturnMain() {
	x, pi := multiReturn(12)
	fmt.Printf("%d, pi: %.2f\n", x, pi)
}
func multiReturn(input uint) (uint, float64) {
	return input, 3.14
}

// 7.3 variadic functions ... ref. page 87
// Java calls varages, variable arguments
func variadicMain() {
	fmt.Println(variadic(3, "a", "b", "d"))

}
func variadic(x int, names ... string) string {
	result := strconv.Itoa(x)
	for _, v := range names {
		result += v
	}
	return result
}

// 7.4 Closure ... ref. page 88 
// create function within a function then access local var
// Java calls annonymous class

// 7.5 Recursion

