package main

import "fmt"

func main() {
	bunchOfNum := []float64{5, 7, 10}
	fmt.Printf("%.2f\n", average(bunchOfNum))
	mainHello()
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

// 7.2 return multiple types ... ref. page 86

// 7.3 variadic functions ... ref. page 87
// Java calls varages, variable arguments

// 7.4 Closure ... ref. page 88 
// create function within a function then access local var
// Java calls annonymous class

// 7.5 Recursion

