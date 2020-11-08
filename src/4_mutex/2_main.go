package main

import (
	"fmt"
	"sync"
)

/**
  Program to illustrate the use of buffered channels of capacity 1 to ensure only one goroutine access the critical section at any point of time
*/

var k int32

func main() {
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Incr(&wg, ch)
	}
	wg.Wait()
	fmt.Print(k)
}

func Incr(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	k += 1
	<-ch
	wg.Done()
}
