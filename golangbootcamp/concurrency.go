package main

import (
	"fmt"
	"time"
)

func getGoroutine() {
	go say("world")
	say("hellow")
}

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//Channel=====================================

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //send sum to c
}

func getChannel() {
	a := []int{1, 2, 3, 4, 5}

	c := make(chan int)
	go sum(a, c)
	x := <-c
	fmt.Println("The channel value is", x)
}

//Buffered Channel =========================

func getBufferedChannel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	c3 := func() { c <- 3 }
	go c3()
	fmt.Println(<-c)
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//1 2 4 3
}

//Range and Close =========================

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func getRangeAndCloseChannel() {
	c := make(chan int) // 两种都可以

	// c := make(chan int, 4)
	// go fibonacci(cap(c), c)
	go fibonacci(4, c)
	fmt.Println("getRangeAndCloseChannel")
	for i := range c {
		fmt.Println(i)
	}

}

//Select =========================
// first come first serve

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for { // infinite loop
		select {
		case c <- x: //2. ready, go!
			x, y = y, x+y
		case <-quit: // 5. receive the quit signal
			fmt.Println("quit")
			return // stop the loop
		}
	}
}

func getSelectGoroutine() {
	fmt.Println("getSelectGoroutine")
	c := make(chan int)
	quit := make(chan int)
	go func() { // 1. start goroutine
		for i := 0; i < 4; i++ {
			fmt.Println(<-c) // wait to receive value, hold here
			// 3. receivet the number
		}
		quit <- 0 // 4. send the quit signal
	}() // IIFE
	fibonacci2(c, quit)

}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "Process done"
}

func getSelectGoroutine2() {
	fmt.Println("getSelectGoroutine2")
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("Received value", v)
			return
		default:
			fmt.Println("Nothing returned")
		}
	}

}
