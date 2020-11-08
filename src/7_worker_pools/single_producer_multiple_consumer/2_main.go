package single_producer_multiple_consumer

import "fmt"

func producer(p chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Produced:", i)
		p <- fmt.Sprintf("%d", i)
	}
	close(p) // you can read from closed channel, writing to a closed channel panics
}

func consumer(i int, p chan string, done chan bool) {
	for m := range p {
		fmt.Printf("consumed: %s by worker %d \n", m, i)

	}
	done <- true
}

func main() {
	done := make(chan bool)
	producerQ := make(chan string)
	go producer(producerQ)
	for i := 1; i < 2; i++ {
		go consumer(i, producerQ, done)
	}
	<-done
	fmt.Println("Main Finished")
}
