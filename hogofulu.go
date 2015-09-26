package main

import (
	"fmt"
	)

// Chapter 10 Concurrency
// Making progress on more than one task simultane-ously is known as concurrency.
// i.e., Eat Bento-box lunch, it has rice, tempula, vegi, sushi.
// eat little bit of each in round robbi(??)
//
//
// Goroutine
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
func mainGoroutine() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}


func main() {
	mainHiva()
	mainGoroutine()
}

// ch 10.2
// Channels:- provide a way for two goroutines to commu-nicate with one another and synchronize their execu-tion.


