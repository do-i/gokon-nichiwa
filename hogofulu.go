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

// select statement
func selectMain() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from one 11111111111"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from two 222222222222"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			// all case gets executed synchronously
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			}
		}
	}()
	fmt.Scanln(new(string))
}

// Buffered Channels (async)
// blocking happens only if spaces are empty or full
func bufferedChannels() {
	// create async channel with 5 spaces
	chenAsync := make(chan int, 5)
	go func() {
		for count := 1; ;count++ {
			chenAsync <- count
			fmt.Println("Producer " + strconv.Itoa(count))
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			data := <- chenAsync
			fmt.Println("Consumer " + strconv.Itoa(data))
			time.Sleep(time.Second * 2)
		}
	}()
	fmt.Scanln(new(string))
}


func main() {
	// mainHiva()
	// mainGoroutine()
	// channelsMain()
	// selectMain()
	bufferedChannels()
}
