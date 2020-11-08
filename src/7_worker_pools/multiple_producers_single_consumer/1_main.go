package main

import (
	"fmt"
	"sync"
)

func producer(w int, p chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("Produced: %d by worker %d \n", i, w)
		p <- fmt.Sprintf("%d", i)
	}
}

func consumer(p <-chan string, done chan bool) {
	for m := range p {
		fmt.Println("consumed:", m)
	}
	done <- true
}

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	producerQ := make(chan string)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go producer(i, producerQ, &wg)
	}
	go consumer(producerQ, done)
	wg.Wait()

	close(producerQ)
	<-done
	fmt.Print("Main Finished")
}
