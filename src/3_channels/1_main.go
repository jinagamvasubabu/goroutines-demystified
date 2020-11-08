package main

import (
	"fmt"
)

/**
  Using Channels we have removed the dependency of Time.sleep instead we have used the blocking nature of channels
  Main Goroutine waits till the value is read from the channel as channels are blocking in nature
*/
func main() {
	ch := make(chan bool)
	go helloWorld("Vasu", ch)
	<-ch
}

func helloWorld(name string, ch chan bool) {
	fmt.Println("Hello World:", name)
	ch <- true
}
