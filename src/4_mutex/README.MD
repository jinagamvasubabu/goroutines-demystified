### Mutex:
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

