package main

import (
	"fmt"
	"runtime"
)

/** unbuffered channels (channels with no capacity) are blocking by nature
* Channel with no capacity
* Sending data to closed channel will PANIC
* Sending to a channel with out receiver will create a leak (orphan goroutine where memory will never be released)
* Receiving from a channel where there is no data in the channel creates deadlock
 */

func main() {
	ch := make(chan string)
	go ComposeEmail(ch)
	go PackBag(ch)
	go LogOffWork(ch)

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("Main goroutine ends %d", runtime.NumGoroutine()) //No receivers available ? then it creates orphan goroutines
}

func ComposeEmail(ch chan string) {
	ch <- "Composed Email"
}

func PackBag(ch chan string) {
	ch <- "Packed Bags"
}

func LogOffWork(ch chan string) {
	ch <- "Logged Off"
}
