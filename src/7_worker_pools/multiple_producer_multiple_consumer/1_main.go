package main

import (
	"fmt"
	"sync"
)

func producer(worker int, p chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("Produced: %d by worker %d \n", i, worker)
		p <- fmt.Sprintf("%d", i)
	}
}

func consumer(i int, p chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for m := range p {
		fmt.Printf("consumed: %s by worker %d \n", m, i)
	}

}

func main() {
	var wp sync.WaitGroup
	var wc sync.WaitGroup
	producerQ := make(chan string)
	for i := 0; i < 3; i++ {
		wp.Add(1)
		go producer(i, producerQ, &wp)
	}

	for i := 1; i < 3; i++ {
		wc.Add(1)
		go consumer(i, producerQ, &wc)
	}
	wp.Wait()
	close(producerQ)
	wc.Wait()
	fmt.Println("Main Finished")
}
