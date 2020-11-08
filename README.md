# goroutines-demystified
A simple and easy explanation of Goroutines (Concurrency&amp;Parallelism) with easy examples

## 1. Concurrency is not parallelism
##### **Parallelism**:
* Doing Multiple things at once
* Number of tasks you are doing in parallel is directly proportional to the number of cores in your system
* Order of execution is not predictable and we should not rely on them to execute in any specific order

Eg:
* Dreaming and sleeping at the same time
* Drinking and Dancing at the same time


##### 2. Concurrency:
* Concurrency is nothing but designing of your code in to an independently executable function and should not have any side effects even if you  run 100,000 times
* Running independently executable function separately and can communicate in between to synchronize their tasks

Eg: Multitasking on single core machine
 
Goroutines
--- 
* Goroutines referred as light weight thread of execution and can run in thousands and its not heavy as 
traditional threads (Referring to Java here)

* Goroutines takes few kb in stack size and the stack can grow or decrease according to the needs of the application.
Where as a thread has to initialized with fixed stack size

* Goroutines are concurrent in nature and uses channels to communicate in between to synchronize instead of creating deadlocks or race conditions

* Any Concurrent programming requires an independently executable function and making that function to a goroutine is very simple

* Goroutine uses Channels to communicate with other goroutines 

* Goroutines uses CSP (Communicating Sequential Processes) to talk to other goroutines by passing data between channels instead of sharing data
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

Because Main is also a goroutine here and its running concurrently and doesn't wait for other goroutines to finish because it 
doesn't know about other goroutines even exists

###### We Need Blocking!!!! Channels and Waitgroups are the way :)


### 3. Channels
 * Channels are used to communicate between processes (goroutines) and its like a two ended pipe
 * Channels send and receive are blocking by nature
 * Data is communicated using channels and no shared data
 
 #### Blocking Breaks Concurrency:
 * Blocking leads to deadlocks
 * Blocking can prevent scaling 
 
 ### Properties:
 * Channel must have a type 
 * A Nil channel is of no use 
 * Sending data to closed channel will PANIC
 * Sending to a channel will block the execution till receiver receives the inputted message, using this behavior we can write concurrent applications
 * Waiting for a Receiver to receive from a channel having values will block (Deadlock)
 * Sending to a channel but there is no receiver will block (Deadlock)
 
 ### UnBuffered Channel:
 * Channel with no capacity
 * Sending data to closed channel will PANIC
 * Sending to a channel with out receiver will create a leak (orphan goroutine where memory will never be released)
 * Receiving from a channel where there is no data in the channel creates deadlock
 
 ### Buffered Channel:
 * Channel with a capacity
 * Sends to a buffered channel are blocked only when the buffer is full
 * Receives from a buffered channel are blocked only when the buffer is empty

### 4. Mutex:
 Before going to Mutex we should have an idea about Critical Sections:
 
 Critical Section is nothing but a section of code which contains mutation of a state of a shared variable 
 
 
 #### Mutex is meant for this reason, at a time only one thread (goroutine) enters into that critical section to avoid any race conditions
 
 We can fix race conditions using both mutexes and buffered Channels of size 1 to ensure that one goroutine access the critical section at any point of time
 ```go
 mutex.Lock()  
 a = a+1 
 mutex.Unlock()  
 ```
 
 ```go
 ch := make(chan bool, 1)
 ch <- true  
 a = a+1 
 <-ch
 ```

### 5. Waitgroups
* Waitgroups is used for waiting for a collection of goroutines to completes their tasks 
* It provides a goroutine synchronization mechanism

```go
wg.Add(i) - Adds i number of gorutines to the wait group 
wg.wait() - wait for all goroutines to finish
```


### 6. Select:
* Select statement blocks until one of the send/receive operation is ready. Takes any random one if both of them ready
* Select statement mainly used with the channels when you dont know which one is ready either send/receive
* Select statement default case is used to avoid panics if none of them ready and avoid blocking




