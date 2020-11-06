Goroutines
--- 
* Goroutines referred as light weight thread of execution and can run in thousands and its not heavy as 
traditional threads (Referring to Java here)

* Goroutines takes few kb in stack size and the stack can grow or decrease according to the needs of the application.
Where as a thread has to initialized with fixed stack size

* Goroutines are concurrent in nature and uses channels to communicate in between to synchronize instead of creating deadlocks or race conditions

* Any Concurrent programming requires an independently executable function and making that function to a goroutine is very simple
```go
go func() {
}() //Anonymous goroutine

go hello() // Add a go keyword before a function creates a concurrently running goroutine
```

Let's write a easy goroutines to understand more about it

```go
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
```

why do you need `time.Sleep(time.Second)` ?
Because Main is also a goroutine here and its concurrently and doesn't wait for other goroutines to finish because it 
doesn't know about other goroutines exist

### We Need Blocking!!!! Channels is the way :)


