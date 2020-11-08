package main

/**
  Program to illustrate the use of mutex to lock on critical section in concurrent world
*/
import (
	"fmt"
	"sync"
)

var j int32
var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Increment(&wg)
	}
	wg.Wait()
	fmt.Print(j)
}

func Increment(wg *sync.WaitGroup) {
	m.Lock()
	j += 1
	m.Unlock()
	wg.Done()
}
