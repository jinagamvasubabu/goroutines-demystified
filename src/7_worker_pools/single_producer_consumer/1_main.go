//single producer and single consumer problem
package main

import "fmt"

//Single Producer single consumer
func producer(p chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Produced:", i)
		p <- fmt.Sprintf("Produced:%d", i)
	}
	close(p) // you can read from closed channel, writing to a closed channel panics
}

func consumer(p chan string, done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("consumed:", <-p)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	producerQ := make(chan string)
	go producer(producerQ)
	go consumer(producerQ, done)
	<-done
	fmt.Print("Main Finished")
}
