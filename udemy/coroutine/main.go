package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://amazon.com",
	}
	c := make(chan string) // channel between coroutine

	for _, link := range links {
		go checkLink(link, c) // start another goroutine
	}

	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<-c) // main routine, wait for the msg to sent in
	// 	go checkLink(<-c, c) // aka link
	// }
	// for {
	// 	// fmt.Println(<-c) // main routine, wait for the msg to sent in
	// 	go checkLink(<-c, c) // aka link
	// }
	//--> same as
	// for v := range c {
	// 	go checkLink(v, c) // aka link
	// }
	for v := range c {
		go func(value string) { // Lambda function
			time.Sleep(5 * time.Second)
			checkLink(value, c)
		}(v)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down ")
		c <- link
		return
	}

	fmt.Println(link, " is up ")
	c <- link
}
