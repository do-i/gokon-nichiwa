package main

import (
	"fmt"
	"time"
	"strconv"
	)

// Chapter 10 Concurrency
// Making progress on more than one task simultane-ously is known as concurrency.
// i.e., Eat Bento-box lunch, it has rice, tempula, vegi, sushi.
// eat little bit of each in round robbi(??)
//
//
// Goroutine
func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println(n, ":", i)
	}
}
func mainGoroutine() {
	for i := 0; i < 3; i++ {
		go f(i)
	}
	var input string
	fmt.Println("hit any key to stop goroutine!!")
	fmt.Scanln(&input)
}



// ch 10.2
// Channels:- provide a way for two goroutines to commu-nicate with one another and synchronize their execu-tion.
// channel direction explicitly specify chan<- as consumer, <-chan as producer.
// implicitly bi-directional
func producer(c chan<- string, producerName string) {
	for i := 0; ; i++ {
		c <- producerName + "ping : " + strconv.Itoa(i)
		time.Sleep(time.Second * 2)
	}
}
func consumer(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
// bi-directional channel : what's the usecase?
func biway(c chan string) {
	for {
		c <- "INPUT"
		booom := <- c
		fmt.Println("output: " + booom)
		time.Sleep(time.Second * 1)
	}
}
func channelsMain() {
	var c chan string = make(chan string)
	go producer(c, "Peoria")
	go producer(c, "Cermak")
	go consumer(c)
	go biway(c)
	var input string
	fmt.Println("hit any key to stop goroutine with channels!!")
	fmt.Scanln(&input)
}


func main() {
	mainHiva()
	mainGoroutine()
	channelsMain()
}