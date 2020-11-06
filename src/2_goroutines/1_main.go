package main

import (
	"fmt"
	"time"
)

func main() {
	go helloWorld("vasu")
	// fmt.Println(runtime.NumGoroutine()) this gives information of how many goroutines exists at that point
	time.Sleep(time.Second) // why do you need this
}

func helloWorld(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("HelloWorld: " + name)
	}
}
