package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go helloWorld("vasu", &wg)
	wg.Wait()
	// fmt.Println(runtime.NumGoroutine()) this gives information of how many goroutines exists at that point
	//time.Sleep(time.Second) // you dont need this now, because waitgroups blocks the main thread till the wait group finishes it tasks
}

func helloWorld(name string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println("HelloWorld: " + name)
	}
	wg.Done()
}
