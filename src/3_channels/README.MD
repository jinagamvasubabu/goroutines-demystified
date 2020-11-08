### Channels
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
